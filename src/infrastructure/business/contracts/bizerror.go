package contracts

import (
	"fmt"
	"strings"
)

type BizError struct {
	Message string
	Status  ResultStatus
	//SubCode int
}

var BizErr *BizError = &BizError{}

func NewBizError(message string, status ResultStatus) *BizError {
	return &BizError{
		Message: message,
		Status:  status,
	}
}

func (er *BizError) Error() string {
	return fmt.Sprintf("Business Error(Status:%s): %s", er.Status, er.Message)
}

func (er *BizError) Is(err error) bool {
	_, ok := err.(*BizError)
	return ok
}

func IsLikeBizError(err error) bool {
	return strings.Contains(err.Error(), "Business Error(Status:")
}
