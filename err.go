package goerr

import "net/http"

type Error struct {
	Code    int
	Message interface{}
	Data    interface{}
}

func ServerError(m ...interface{}) *Error {
	var message interface{}
	var data interface{}
	message = "Something went wrong. Please try again."

	if len(m) > 0 {
		data = m[0]
	}
	if len(m) > 1 {
		message = m[1]
	}
	return &Error{
		Code:    http.StatusInternalServerError,
		Message: message,
		Data:    data,
	}
}

func NoRecordFound(m ...interface{}) *Error {
	var message interface{}
	var data interface{}
	message = "Record not found"
	if len(m) > 0 {
		data = m[0]
	}
	if len(m) > 1 {
		message = m[1]
	}
	return &Error{
		Code:    http.StatusNotFound,
		Message: message,
		Data:    data,
	}
}

func BadRequest(m ...interface{}) *Error {
	var message interface{}
	var data interface{}
	message = "The given data was invalid"

	if len(m) > 0 {
		data = m[0]
	}
	if len(m) > 1 {
		message = m[1]
	}
	return &Error{
		Code:    http.StatusBadRequest,
		Message: message,
		Data:    data,
	}
}
