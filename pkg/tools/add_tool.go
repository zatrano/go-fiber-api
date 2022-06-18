package tools

import (
	"github.com/gofiber/fiber/v2"
	"github.com/zatrano/go-fiber-api/common/models"
)

type AddToolRequestBody struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

func (h handler) AddTool(c *fiber.Ctx) error {
	body := AddToolRequestBody{}

	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	var tool models.Tool

	tool.Title = body.Title
	tool.Description = body.Description

	if result := h.DB.Create(&tool); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(&tool)
}
