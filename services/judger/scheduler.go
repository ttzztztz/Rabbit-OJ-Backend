package judger

import (
	"Rabbit-OJ-Backend/models"
	"Rabbit-OJ-Backend/protobuf"
	"Rabbit-OJ-Backend/utils"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"strconv"
)

type CollectedStdout struct {
	Stdout      string
	RightStdout string
}

func Scheduler(request *protobuf.JudgeRequest) error {
	sid := request.Sid

	fmt.Println("[Scheduler] Received judge request " + sid)

	// init path
	currentPath, err := utils.JudgeGenerateDirWithMkdir(sid)
	if err != nil {
		return err
	}

	outputPath, err := utils.JudgeGenerateOutputDirWithMkdir(currentPath)
	if err != nil {
		return err
	}

	codePath := fmt.Sprintf("%s/%s.code", currentPath, sid)
	if err := ioutil.WriteFile(codePath, request.Code, 0644); err != nil {
		return err
	}

	jsonPath := codePath + ".json"
	casePath, err := utils.JudgeCaseDir(request.Tid, request.Version)
	if err != nil {
		return err
	}

	compileInfo, ok := utils.CompileObject[request.Language]
	if !ok {
		return errors.New("language doesn't support")
	}

	// get case
	storage, err := InitTestCase(request.Tid, request.Version)
	if err != nil {
		return err
	}

	// compile
	if err := Compiler(codePath, &compileInfo); err != nil {
		fmt.Println("CE", err)
		return callbackAllError("CE", sid, storage)
	}

	// run
	if err := Runner(
		codePath,
		&compileInfo,
		strconv.FormatUint(uint64(storage.DatasetCount), 10),
		strconv.FormatUint(uint64(request.TimeLimit), 10),
		strconv.FormatUint(uint64(request.SpaceLimit), 10),
		casePath,
		outputPath); err != nil {

		fmt.Println("RE", err)
		return callbackAllError("RE", sid, storage)
	}

	jsonFileByte, err := ioutil.ReadFile(jsonPath)
	if err != nil {
		return callbackAllError("RE", sid, storage)
	}

	var testResultArr []models.TestResult
	if err := json.Unmarshal(jsonFileByte, &testResultArr); err != nil {
		return callbackAllError("RE", sid, storage)
	}

	// collect std::out
	fmt.Println("Collecting stdout " + sid)
	allStdin := make([]CollectedStdout, storage.DatasetCount)
	for i := uint32(0); i < storage.DatasetCount; i++ {

		path, err := utils.JudgeFilePath(
			storage.Tid,
			storage.Version,
			strconv.FormatUint(uint64(i), 10),
			"out")

		if err != nil {
			return err
		}

		stdoutByte, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}

		allStdin[i].RightStdout = string(stdoutByte)
	}

	for i := uint32(0); i < storage.DatasetCount; i++ {
		path := fmt.Sprintf("%s/%d.out", outputPath, i)

		stdoutByte, err := ioutil.ReadFile(path)
		if err != nil {
			allStdin[i].Stdout = ""
		} else {
			allStdin[i].Stdout = string(stdoutByte)
		}
	}
	// judge std::out
	fmt.Println("Judging stdout " + sid)
	resultList := make([]*protobuf.JudgeCaseResult, storage.DatasetCount)

	for index, item := range allStdin {
		testResult := &testResultArr[index]

		judgeResult := JudgeOneCase(testResult, item.Stdout, item.RightStdout, request.CompMode)

		resultList[index].Status = judgeResult.Status
		resultList[index].SpaceUsed = judgeResult.SpaceUsed
		resultList[index].TimeUsed = judgeResult.TimeUsed
	}
	// mq return result
	go callbackWebSocket(sid)
	if err := callbackSuccess(
		sid,
		resultList); err != nil {
		return err
	}
	// todo: clear cache

	return nil
}