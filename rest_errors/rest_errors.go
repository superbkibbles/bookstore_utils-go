package rest_errors

import (
	"errors"
	"net/http"
)

type RestErr struct {
	Message string        `json:"message"`
	Status  int           `json:"status"`
	Error   string        `json:"error"`
	Causes  []interface{} `json:"causes"`
}

func NewError(msg string) error {
	return errors.New(msg)
}

func NewBadRequestErr(message string) *RestErr {
	return &RestErr{
		Message: message,
		Error:   "Bad_request",
		Status:  http.StatusBadRequest,
	}
}

func NewNotFoundErr(message string) *RestErr {
	return &RestErr{
		Message: message,
		Error:   "Not_found",
		Status:  http.StatusNotFound,
	}
}

func NewInternalServerErr(message string, err error) *RestErr {
	result := &RestErr{
		Message: message,
		Error:   "internal_server_error",
		Status:  http.StatusInternalServerError,
	}
	if err != nil {
		result.Causes = []interface{}{err.Error()}
	}
	return result
}
