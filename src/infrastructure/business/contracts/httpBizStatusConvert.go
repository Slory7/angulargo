package contracts

import (
	"net/http"
)

func HttpToBizStatus(httpStatus int) (result ResultStatus) {
	result = Success
	switch {
	case httpStatus == http.StatusNotFound:
		result = NotFound
	case httpStatus == http.StatusConflict:
		result = Conflict
	case httpStatus == http.StatusForbidden:
		result = Forbidden
	case httpStatus == http.StatusUnprocessableEntity:
		result = BadLogic
	case httpStatus == http.StatusBadRequest:
		result = BadData
	case httpStatus == http.StatusUnauthorized:
		result = Unauthorized
	case httpStatus == http.StatusRequestTimeout:
		result = Timeout
	case httpStatus >= http.StatusBadRequest:
		result = Error
	}
	return
}

func BizStatusToHttp(status ResultStatus) (result int) {
	result = http.StatusOK
	switch {
	case status == NotFound:
		result = http.StatusNotFound
	case status == Conflict:
		result = http.StatusConflict
	case status == Forbidden:
		result = http.StatusForbidden
	case status == BadLogic:
		result = http.StatusUnprocessableEntity
	case status == BadData:
		result = http.StatusBadRequest
	case status == Unauthorized:
		result = http.StatusUnauthorized
	case status == Timeout:
		result = http.StatusRequestTimeout
	case status == Error:
		result = http.StatusInternalServerError
	}
	return
}
