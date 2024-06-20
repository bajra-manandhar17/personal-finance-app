package excep

import (
	"log"
	"net/http"
)

type HttpExcep struct {
	Status  int    `json:"status"`
	Type    string `json:"type"`
	Details string `json:"details"`
}

var exceptionTypeToHttpStatus = map[string]int{
	EXCEP_USER_NOT_FOUND:        http.StatusNotFound,
	EXCEP_EMAIL_EXISTS:          http.StatusConflict,
	EXCEP_UNAUTHORIZED:          http.StatusUnauthorized,
	EXCEP_INVALID_PAYLOAD:       http.StatusBadRequest,
	EXCEP_FORBIDDEN:             http.StatusForbidden,
	EXCEP_INTERNAL_SERVER_ERROR: http.StatusInternalServerError,
}

func MapErrorToHttpException(err error) HttpExcep {
	exceptionType := EXCEP_INTERNAL_SERVER_ERROR
	details := "Internal Server Error"

	if exceptionIns, ok := err.(DomainExcep); ok {
		exceptionType = exceptionIns.Type
		details = exceptionIns.Details
	} else {
		log.Printf("Internal Server Error: %v", err.Error())
	}

	return HttpExcep{
		Status:  getHttpStatusForExceptionType(exceptionType),
		Type:    exceptionType,
		Details: details,
	}
}

func getHttpStatusForExceptionType(exceptionType string) int {
	if status, ok := exceptionTypeToHttpStatus[exceptionType]; ok {
		return status
	}
	return http.StatusInternalServerError
}
