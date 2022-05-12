package errs

import (
	"github.com/gofiber/fiber/v2"
)

type AppError struct {
	Code        int
	Message     string
	MessageCode int
}

func (e AppError) Error() string {
	return e.Message
}

func NewNotFoundError(message string) error {
	return AppError{
		Code:    fiber.StatusNotFound,
		Message: message,
	}
}

func NewUnexpectedError() error {
	return AppError{
		Code:    fiber.StatusInternalServerError,
		Message: "unexpected error",
	}
}

func NewValidationError(message string) error {
	return AppError{
		Code:    fiber.StatusUnprocessableEntity,
		Message: message,
	}
}

func CustomError(message string, messagecode int) error {
	return AppError{
		Code:        fiber.StatusBadRequest,
		Message:     message,
		MessageCode: messagecode,
	}
}
