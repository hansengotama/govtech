package errcustom

import (
	"fmt"
	"net/http"
)

type ErrorInvalidField struct {
	Field         string
	AcceptedValue string
}

func (err ErrorInvalidField) Error() string {
	return fmt.Sprintf("The value of field %s is invalid. Accepted value %s", err.Field, err.AcceptedValue)
}

type ErrorFieldLessThanMinValue struct {
	Field    string
	MinValue int
}

func (err ErrorFieldLessThanMinValue) Error() string {
	return fmt.Sprintf("The value of field %s must be greater than or equal to %v", err.Field, err.MinValue)
}

type ErrorFieldMoreThanMaxValue struct {
	Field    string
	MaxValue int
}

func (err ErrorFieldMoreThanMaxValue) Error() string {
	return fmt.Sprintf("The value of field %s must be less than or equal to %v", err.Field, err.MaxValue)
}

func GetHTTPErrCode(err error) int {
	switch errCustom := err.(type) {
	case ErrorFieldLessThanMinValue:
		return http.StatusBadRequest
	case ErrorFieldMoreThanMaxValue:
		return http.StatusBadRequest
	case ErrorInvalidField:
		return http.StatusBadRequest
	default:
		fmt.Println(errCustom)
		return http.StatusInternalServerError
	}
}
