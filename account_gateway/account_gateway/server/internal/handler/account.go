package handler

import (
	"identity_provider/internal/service"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

type accountHandler struct {
	accountSrv service.AccountService
}

func NewAccountHandler(accountSrv service.AccountService) accountHandler {
	return accountHandler{accountSrv: accountSrv}
}

func (h accountHandler) CreateAccount(c *fiber.Ctx) error {
	c.Accepts("application/json")

	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	iss := claims["iss"].(string)
	userid, _ := strconv.Atoi(iss)

	body := new(service.CreateAccountForm)
	if err := c.BodyParser(body); err != nil {
		return fiber.ErrBadRequest
	}
	body.UserID = int(userid)

	result, err := h.accountSrv.CreateAccount(body)
	if err != nil {
		return handleError(c, err)
	}
	return c.Status(http.StatusOK).JSON(&fiber.Map{
		"data": result,
	})

}

func (h accountHandler) FindAccount(c *fiber.Ctx) error {

	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	iss := claims["iss"].(string)
	userid, _ := strconv.Atoi(iss)
	body := service.RequstFindAccount{
		UserID: int(userid),
	}

	result, err := h.accountSrv.FindAccount(&body)
	if err != nil {
		return handleError(c, err)
	}
	return c.Status(http.StatusOK).JSON(&fiber.Map{
		"data": result,
	})

}

func (h accountHandler) UpdateAccount(c *fiber.Ctx) error {
	c.Accepts("application/json")

	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	iss := claims["iss"].(string)
	userid, _ := strconv.Atoi(iss)

	body := new(service.UpdateAccountForm)

	if err := c.BodyParser(body); err != nil {
		return fiber.ErrBadRequest
	}
	body.UserID = int(userid)

	result, err := h.accountSrv.UpdateAccount(body)
	if err != nil {
		return handleError(c, err)
	}
	return c.Status(http.StatusOK).JSON(&fiber.Map{
		"data": result,
	})

}
