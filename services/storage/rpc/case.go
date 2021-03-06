package rpc

import (
	"Rabbit-OJ-Backend/models/protobuf"
	"Rabbit-OJ-Backend/services/question"
	"fmt"
)

type CaseService struct{}

func (s *CaseService) Case(request protobuf.TestCaseRequest, response *protobuf.TestCaseResponse) error {
	tid := request.Tid

	fmt.Printf("[RPC] test case received request %d \n", tid)
	if err := question.Case(tid, response); err != nil {
		return err
	}

	return nil
}
