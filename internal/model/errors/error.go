package errors

import "net/http"

type Error struct {
	Code int
	Msg  string
}

const ErrResourceNotFound = "not found"

func New(code int, msg string) *Error {
	return &Error{
		Code: code,
		Msg:  msg,
	}
}

func NewInternalServerError500(msg string) *Error {
	return New(http.StatusInternalServerError, msg)
}

func NewResourceNotFound404(msg string) *Error {
	return New(http.StatusNotFound, msg)
}

func NewBadRequest400(msg string) *Error {
	return New(http.StatusBadRequest, msg)
}
