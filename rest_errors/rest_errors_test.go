package rest_errors

import (
	"errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewInternalServerError(t *testing.T) {
	err := NewInternalServerErr("this is message", errors.New("database error"))

	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "this is message", err.Message)
	assert.EqualValues(t, "internal_server_error", err.Error)

	assert.NotNil(t, err.Causes)
	assert.EqualValues(t, 1, len(err.Causes))
	assert.EqualValues(t, "database error", err.Causes[0])
}

func TestNewBadRequestError(t *testing.T) {
	err := NewBadRequestErr("this is message")
	assert.NotNil(t, err)
	assert.EqualValues(t, "this is message", err.Message)
	assert.EqualValues(t, "Bad_request", err.Error)
}

func TestNewNotFoundError(t *testing.T) {
	err := NewNotFoundErr("this is message")
	assert.NotNil(t, err)
	assert.EqualValues(t, "this is message", err.Message)
	assert.EqualValues(t, "Not_found", err.Error)
}

func TestNewError(t *testing.T) {
	err := NewError("message")
	assert.NotNil(t, err)
	assert.EqualValues(t, "message", err.Error())
}