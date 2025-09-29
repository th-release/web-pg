package auth

import (
	"cth.release/web-pg/common"
	"github.com/gofiber/fiber/v2"
)

func GetKey(c *fiber.Ctx) error {
	var req GetKeyDto
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(common.BasicResponse{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	token, err := common.CreateAuthKey(req.Database, req.Username, req.Password)

	if err != nil {
		return c.Status(500).JSON(common.BasicResponse{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	return c.Status(201).JSON(common.BasicResponse{
		Success: true,
		Message: "",
		Data:    token,
	})
}
