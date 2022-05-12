package handler

import (
	"identity_provider/internal/service"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type authHandler struct {
	authSrv service.AuthService
}

func NewAuthHandler(authSrv service.AuthService) authHandler {
	return authHandler{authSrv: authSrv}
}

func (h authHandler) DefaultLogin(c *fiber.Ctx) error {
	c.Accepts("application/json")
	body := new(service.LoginDefaultForm)
	if err := c.BodyParser(body); err != nil {
		return fiber.ErrBadRequest
	}
	result, err := h.authSrv.DefaultLogin(body)
	if err != nil {
		return handleError(c, err)
	}
	return c.Status(http.StatusOK).JSON(&fiber.Map{
		"data": result,
	})

}
