package handler

import (
	"identity_provider/internal/service"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"

	"github.com/golang-jwt/jwt/v4"
)

type userHandler struct {
	userSrv service.UserService
}

func NewUserHandler(userSrv service.UserService) userHandler {
	return userHandler{userSrv: userSrv}
}

func (h userHandler) GetUsers(c *fiber.Ctx) error {
	users, err := h.userSrv.GetAllUsers()
	if err != nil {
		return err
	}
	return c.Status(http.StatusOK).JSON(&fiber.Map{
		"success": true,
		"users":   users,
	})

}

func (h userHandler) CreateUser(c *fiber.Ctx) error {
	c.Accepts("application/json")
	body := new(service.UserRegisterForm)
	if err := c.BodyParser(body); err != nil {
		return fiber.ErrBadRequest
	}
	result, err := h.userSrv.CreateUser(body)
	if err != nil {
		return handleError(c, err)
	}
	return c.Status(http.StatusOK).JSON(&fiber.Map{
		"data": result,
	})

}

func (h userHandler) AcceptTermCondition(c *fiber.Ctx) error {
	c.Accepts("application/json")

	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	iss := claims["iss"].(string)
	userid, _ := strconv.Atoi(iss)

	body := new(service.UserAccpetTermConditionForm)
	if err := c.BodyParser(body); err != nil {
		return fiber.ErrBadRequest
	}
	body.UserID = uint(userid)

	result, err := h.userSrv.AcceptTermCondition(body)
	if err != nil {
		return handleError(c, err)
	}
	return c.Status(http.StatusOK).JSON(&fiber.Map{
		"data": result,
	})

}
