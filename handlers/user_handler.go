package handlers

import (
	"github.com/Jere283/autek_rest_api/service"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	UserService service.UserService
}

func (h *UserHandler) GetUser(c *fiber.Ctx) error {
	id := c.Params("id")
	user, err := h.UserService.GetUserByID(id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "User not found", "data": nil})
	}
	return c.JSON(fiber.Map{"status": "success", "message": "User found", "data": user})
}
