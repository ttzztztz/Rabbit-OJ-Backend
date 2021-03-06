package initialize

import (
	StorageService "Rabbit-OJ-Backend/services/storage"
	"Rabbit-OJ-Backend/utils/files"
	"strconv"
)

func CheckTestCase() {
	storage := StorageService.ReadStorageFile()
	newStorage := make([]StorageService.Storage, 0)

	for _, item := range storage {
		valid := true

		for i := uint32(1); i <= item.DatasetCount; i++ {
			inputFilePath, err := files.JudgeFilePath(item.Tid, item.Version, strconv.FormatUint(uint64(i), 10), "in")
			if err != nil {
				valid = false
			}
			if !files.Exists(inputFilePath) {
				valid = false
			}

			outputFilePath, err := files.JudgeFilePath(item.Tid, item.Version, strconv.FormatUint(uint64(i), 10), "out")
			if err != nil {
				valid = false
			}
			if !files.Exists(outputFilePath) {
				valid = false
			}
		}

		if valid {
			newStorage = append(newStorage, item)
		}
	}

	if err := StorageService.SaveStorageFile(newStorage); err != nil {
		panic(err)
	}
}

func PruneTestCase() {
	// todo: delete not referenced object
}
