package acaseSdk

import (
	"github.com/tmconsulting/acase-golang-sdk/acaseSts"
)

type AcaseResponseError struct {
	Message string
	Code    string
}

func (ar *AcaseResponseError) Error() string {
	return ar.Message
}

func (ar *AcaseResponseError) ErrorCode() string {
	return ar.Code
}

func ResponseError(item acaseSts.BaseResponse) *AcaseResponseError {
	if item.Error != nil {
		return &AcaseResponseError{
			Code:    item.Error.Code,
			Message: item.Error.Description,
		}
	}
	return nil
}
