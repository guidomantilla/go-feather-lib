package rest

import (
	"net/http"
)

type Exception struct {
	Code    int      `json:"code,omitempty"`
	Message string   `json:"message,omitempty"`
	Errors  []string `json:"errors,omitempty"`
}

func Convert2StringArray(err ...error) []string {

	var errors []string
	if len(err) >= 0 {
		for _, err2 := range err {
			errors = append(errors, err2.Error())
		}
	}
	return errors
}

func BadRequestException(message string, err ...error) *Exception {

	return &Exception{
		Message: message,
		Errors:  Convert2StringArray(err...),
		Code:    http.StatusBadRequest,
	}
}

func UnauthorizedException(message string, err ...error) *Exception {

	return &Exception{
		Message: message,
		Errors:  Convert2StringArray(err...),
		Code:    http.StatusUnauthorized,
	}
}

func ForbiddenException(message string, err ...error) *Exception {

	return &Exception{
		Message: message,
		Errors:  Convert2StringArray(err...),
		Code:    http.StatusForbidden,
	}
}

func NotFoundException(message string, err ...error) *Exception {

	return &Exception{
		Message: message,
		Errors:  Convert2StringArray(err...),
		Code:    http.StatusNotFound,
	}
}

func InternalServerErrorException(message string, err ...error) *Exception {

	return &Exception{
		Message: message,
		Errors:  Convert2StringArray(err...),
		Code:    http.StatusInternalServerError,
	}
}
