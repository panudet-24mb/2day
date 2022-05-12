package handler

import (
	"identity_provider/internal/errs"

	"github.com/gofiber/fiber/v2"
)

func handleError(c *fiber.Ctx, err error) error {

	switch e := err.(type) {
	case errs.AppError:
		return c.Status(e.Code).JSON(&fiber.Map{
			"data": e.Message,
			"code": e.MessageCode,
		})

	case error:
		return fiber.NewError(fiber.StatusInternalServerError)
	}

	return fiber.NewError(fiber.StatusInternalServerError)

}
