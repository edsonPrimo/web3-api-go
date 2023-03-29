package errors

import (
	"log"
	"strings"

	goerrors "github.com/go-errors/errors"
)

type HttpError struct {
	StatusCode int    `json:"statusCode"`
	ErrorName  string `json:"error"`
	Message    string `json:"message"`
}

func NewError(code int, err error) *HttpError {
	httpError := &HttpError{
		StatusCode: code,
		ErrorName:  "Error",
		Message:    err.Error(),
	}
	if strings.Contains(err.Error(), "not found") {
		httpError.ErrorName = "Not Found"
		httpError.StatusCode = 404
	}
	log.Println(goerrors.New(err).ErrorStack())
	return httpError
}
func NewErrorMessage(message string) *HttpError {
	httpError := &HttpError{
		StatusCode: 400,
		ErrorName:  "Error",
		Message:    message,
	}
	if strings.Contains(message, "not found") {
		httpError.ErrorName = "Not Found"
		httpError.StatusCode = 404
	}
	log.Println(goerrors.Errorf(message))
	return httpError
}
func NewHttpError(code int, name string, message string) *HttpError {
	httpError := &HttpError{
		StatusCode: code,
		ErrorName:  name,
		Message:    message,
	}
	log.Println(goerrors.Errorf(message))
	return httpError
}
func NewNotFoundHttpError(message string) *HttpError {
	log.Println(goerrors.Errorf(message))
	return NewHttpError(404, "Not found", message)
}
func NewInternalErrorHttpError(message string) *HttpError {
	log.Println(goerrors.Errorf(message))
	return NewHttpError(500, "Internal error", message)
}
func NewBadRequestHttpError(message string) *HttpError {
	log.Println(goerrors.Errorf(message))
	return NewHttpError(400, "Bad request", message)
}

func (e *HttpError) Error() string {
	return e.ErrorName
}

func (e *HttpError) Unwrap() error {
	return e
}

func Is(err error, target error) bool {
	return err.Error() == target.Error()
}
