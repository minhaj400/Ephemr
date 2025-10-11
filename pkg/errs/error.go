package errs

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type AppError struct {
	Code       string
	Message    string
	Internal   error
	HTTPStatus int
}

func (e *AppError) Error() string {
	if e.Internal != nil {
		return fmt.Sprintf("[%s] %s | internal: %v", e.Code, e.Message, e.Internal)
	}
	return fmt.Sprintf("[%s] %s", e.Code, e.Message)
}

func New(code, message string, status int, internal error) *AppError {
	return &AppError{
		Code:       code,
		Message:    message,
		HTTPStatus: status,
		Internal:   internal,
	}
}

func From(err error) (*AppError, bool) {
	var appErr *AppError
	if errors.As(err, &appErr) {
		return appErr, true
	}
	return nil, false
}

func InternalError(internal error) *AppError {
	return New("INTERNAL_SERVER_ERROR", "Something went wrong", http.StatusInternalServerError, internal)
}

func NotFound(message string, internal error) *AppError {
	return New("NOT_FOUND", message, http.StatusNotFound, internal)
}

func BadRequest(message string, internal error) *AppError {
	return New("BAD_REQUEST", message, http.StatusBadRequest, internal)
}

func FromValidation(err error) *AppError {
	if verrs, ok := err.(validator.ValidationErrors); ok && len(verrs) > 0 {
		vErr := verrs[0]
		field := vErr.Field()
		tag := vErr.Tag()

		msg := fmt.Sprintf("%s is %s", field, tag)
		return New(ValidationError, msg, http.StatusBadRequest, nil)
	}

	return BadRequest("Invalid request payload", err)
}
