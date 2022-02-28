package common

import "net/http"

type IAppError interface {
	error
	StatusCode() int
}

type AppError struct {
	HttpStatus int    `json:"-"`
	Message    string `json:"message"`
	Cause      error  `json:"-"`
}

func (appError AppError) Error() string {
	return appError.Message
}

func (appError AppError) StatusCode() int {
	return appError.HttpStatus
}

func (appError AppError) WithCause(err error) *AppError {
	appError.Cause = err
	return &appError
}

func BadRequest(message string) *AppError {
	return &AppError{HttpStatus: http.StatusBadRequest, Message: message}
}

func ServerError(message string) *AppError {
	return &AppError{HttpStatus: http.StatusInternalServerError, Message: message}
}
