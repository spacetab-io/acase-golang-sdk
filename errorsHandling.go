package acase_sdk

import (
	"log"
)

type AcaseResponseError struct {
	Message	string
	Code	string
}

func (ar *AcaseResponseError) Error() string {
	return ar.Message
}

func (ar *AcaseResponseError) ErrorCode() string {
	return ar.Code
}

func FatalError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func ErrorResponse(acaseErrors *[]RespError) *[]AcaseResponseError {
	if acaseErrors == nil || len(*acaseErrors) == 0 {
		res := make([]AcaseResponseError, 1)
		res[0].Message = "Undefined error"
		return &res
	}
	res := make([]AcaseResponseError, len(*acaseErrors))
	for i, item := range *acaseErrors {
		res[i].Message = item.Message
		res[i].Code = item.Code
	}
	return &res
}

