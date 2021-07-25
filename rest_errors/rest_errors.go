// package rest_errors
//
// import (
// 	"errors"
// 	"net/http"
// )
//
// type RestErr interface {
// 	Message() string
// 	Status() int
// 	Error() string
// 	Causes() []interface{}
// }
//
// type restErr struct {
// 	message string        `json:"message"`
// 	status  int           `json:"status"`
// 	error   string        `json:"error"`
// 	causes  []interface{} `json:"causes"`
// }
//
// // func (e restErr) Error() string {
// // 	return fmt.Sprintf("message %s- status: %d - error: %s -casues: [ %v ]", e.message, e.status, e.error, e.causes)
// // }
//
// func (e restErr) Error() string {
// 	return e.error
// }
//
// func (e restErr) Message() string {
// 	return e.message
// }
// func (e restErr) Status() int {
// 	return e.status
// }
//
// func (e restErr) Causes() []interface{} {
// 	return e.causes
// }
//
// func NewError(msg string) error {
// 	return errors.New(msg)
// }
//
// func NewRestError(message string, status int, error string, causes []interface{}) RestErr {
// 	return restErr{
// 		message: message,
// 		status:  status,
// 		error:   error,
// 		causes:  causes,
// 	}
// }
//
// func NewBadRequestErr(message string) RestErr {
// 	return restErr{
// 		message: message,
// 		error:   "Bad_request",
// 		status:  http.StatusBadRequest,
// 	}
// }
//
// func NewNotFoundErr(message string) RestErr {
// 	return restErr{
// 		message: message,
// 		error:   "Not_found",
// 		status:  http.StatusNotFound,
// 	}
// }
//
// func NewUnauthorizedError(message string) RestErr {
// 	return restErr{
// 		message: message,
// 		error:   "Unauthirized",
// 		status:  http.StatusUnauthorized,
// 	}
// }
//
// func NewInternalServerErr(message string, err error) RestErr {
// 	result := restErr{
// 		message: message,
// 		error:   "internal_server_error",
// 		status:  http.StatusInternalServerError,
// 	}
// 	if err != nil {
// 		result.causes = append(result.causes, err.Error())
// 	}
// 	return result
// }

package rest_errors

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type RestErr interface {
	Message() string
	Status() int
	ErrorM() string
	Causes() []interface{}
}

type restErr struct {
	ErrMessage string        `json:"message"`
	ErrStatus  int           `json:"status"`
	ErrError   string        `json:"error"`
	ErrCauses  []interface{} `json:"causes"`
}

func (e restErr) ErrorM() string {
	return fmt.Sprintf("message: %s - status: %d - error: %s - causes: %v",
		e.ErrMessage, e.ErrStatus, e.ErrError, e.ErrCauses)
}

func (e restErr) Message() string {
	return e.ErrMessage
}

func (e restErr) Status() int {
	return e.ErrStatus
}

func (e restErr) Causes() []interface{} {
	return e.ErrCauses
}

func NewRestError(message string, status int, err string, causes []interface{}) RestErr {
	return restErr{
		ErrMessage: message,
		ErrStatus:  status,
		ErrError:   err,
		ErrCauses:  causes,
	}
}

func NewRestErrorFromBytes(bytes []byte) (RestErr, error) {
	var apiErr restErr
	if err := json.Unmarshal(bytes, &apiErr); err != nil {
		return nil, errors.New("invalid json")
	}
	return apiErr, nil
}

func NewBadRequestErr(message string) RestErr {
	return restErr{
		ErrMessage: message,
		ErrStatus:  http.StatusBadRequest,
		ErrError:   "bad_request",
	}
}

func NewNotFoundErr(message string) RestErr {
	return restErr{
		ErrMessage: message,
		ErrStatus:  http.StatusNotFound,
		ErrError:   "not_found",
	}
}

func NewUnauthorizedError(message string) RestErr {
	return restErr{
		ErrMessage: message,
		ErrStatus:  http.StatusUnauthorized,
		ErrError:   "unauthorized",
	}
}

func NewInternalServerErr(message string, err error) RestErr {
	result := restErr{
		ErrMessage: message,
		ErrStatus:  http.StatusInternalServerError,
		ErrError:   "internal_server_error",
	}
	if err != nil {
		result.ErrCauses = append(result.ErrCauses, err.Error())
	}
	return result
}
