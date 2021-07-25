package rest_errors

import (
	"errors"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewInternalServerError(t *testing.T) {
	err := NewInternalServerErr("this is message", errors.New("database error"))
	fmt.Println(err)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status())
	assert.EqualValues(t, "this is message", err.Message())
	assert.EqualValues(t, "message: this is message - status: 500 - error: internal_server_error - causes: [database error]", err.ErrorM())

	assert.NotNil(t, err.Causes)
	assert.EqualValues(t, 1, len(err.Causes()))
	assert.EqualValues(t, "database error", err.Causes()[0])
}

// func TestNewBadRequestError(t *testing.T) {
// 	err := NewBadRequestErr("this is message")
// 	assert.NotNil(t, err)
// 	assert.EqualValues(t, "this is message", err.Message)
// 	assert.EqualValues(t, "Bad_request", err.Error)
// }
//
// func TestNewNotFoundError(t *testing.T) {
// 	err := NewNotFoundErr("this is message")
// 	assert.NotNil(t, err)
// 	assert.EqualValues(t, "this is message", err.Message)
// 	assert.EqualValues(t, "Not_found", err.Error)
// }
//
// func TestNewError(t *testing.T) {
// 	err := NewError("message")
// 	assert.NotNil(t, err)
// 	assert.EqualValues(t, "message", err.Error())
// }
//
// func TestNewRestError(t *testing.T) {
// 	errInterface := []interface{}{
// 		"sttrangeErr",
// 	}
// 	err := NewRestError("message", 352, "error_message", errInterface)
// 	assert.NotNil(t, err)
// 	assert.EqualValues(t, "message", err.Message)
// 	assert.EqualValues(t, 352, err.Status)
// 	assert.EqualValues(t, "error_message", err.Error)
// 	assert.EqualValues(t, 1, len(err.Causes))
// 	assert.EqualValues(t, "sttrangeErr", err.Causes[0])
// }
