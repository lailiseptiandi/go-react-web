package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lailiseptiandi/go-web-app/app/dtos"
	"github.com/lailiseptiandi/go-web-app/app/services"
)

type UserHandler struct {
	userService services.UserService
}

func NewUserHandler(s services.UserService) UserHandler {
	return UserHandler{s}
}

func (h *UserHandler) GetUsers(c *fiber.Ctx) error {
	users, err := h.userService.GetAllUsers()
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.Render("user/index", fiber.Map{
		"users": users,
	})
}

func (h *UserHandler) CreateUser(c *fiber.Ctx) error {
	var user dtos.UserCreate
	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	_, err := h.userService.CreateUser(user)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	return c.Redirect("/user", fiber.StatusFound)
}
