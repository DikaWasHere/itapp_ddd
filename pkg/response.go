package pkg

import "github.com/gofiber/fiber/v2"

func SuccessResponse(c *fiber.Ctx, data interface{}) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":    data,
		"message": "success",
		"status":  fiber.StatusOK,
	})
}

func BadRequestResponse(c *fiber.Ctx, message string) error {
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"data":    nil,
		"message": message,
		"status":  fiber.StatusBadRequest,
	})
}
