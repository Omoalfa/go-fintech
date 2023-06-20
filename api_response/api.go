package api_response

import (
	"github.com/gofiber/fiber/v2"
)

func BadRequest(c *fiber.Ctx, errors any) error {
	return c.Status(fiber.StatusBadRequest).JSON(map[string]any{
		"status":  fiber.StatusBadRequest,
		"message": "Bad Request",
		"errors":  errors,
	})
}

func SuccessAction(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(map[string]any{
		"status":  fiber.StatusOK,
		"message": "Success",
		"data":    nil,
	})
}

func Success(c *fiber.Ctx, data any) error {
	return c.Status(fiber.StatusOK).JSON(map[string]any{
		"status":  fiber.StatusOK,
		"message": "Success",
		"data":    data,
	})
}

func SuccessCreated(c *fiber.Ctx, data any) error {
	return c.Status(fiber.StatusCreated).JSON(map[string]any{
		"status":  fiber.StatusCreated,
		"message": "Success",
		"data":    data,
	})
}

func ServerError(c *fiber.Ctx) error {
	return c.Status(fiber.StatusInternalServerError).JSON(map[string]any{
		"status":  fiber.StatusInternalServerError,
		"message": "Something went wrong",
		"data":    nil,
	})
}
