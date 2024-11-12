package pkg

import (
	"errors"
	"github.com/gofiber/fiber/v2"
)

func NewErrorHandler(ctx *fiber.Ctx, err error) error {
	// Init Logger
	logger := NewLogger()

	// Status code defaults to 500
	code := fiber.StatusInternalServerError

	// Retrieve the custom status code if it's a *fiber.Error
	var e *fiber.Error
	if errors.As(err, &e) {
		code = e.Code
	}

	// Return if Not Found
	if code == fiber.StatusNotFound {
		return ctx.Status(code).JSON(fiber.Map{
			"success": false,
			"message": "Not Found",
		})
	}

	// Return if Bad Request
	if code == fiber.StatusBadRequest {
		return ctx.Status(code).JSON(fiber.Map{
			"success":    false,
			"message":    "Bad Request",
			"error_code": "BAD_REQUEST",
			"error":      err.Error(),
		})
	}

	// Return if Unauthorized
	if code == fiber.StatusUnauthorized {
		return ctx.Status(code).JSON(fiber.Map{
			"success": false,
			"message": "Unauthorized",
		})
	}

	// Return if Forbidden
	if code == fiber.StatusForbidden {
		return ctx.Status(code).JSON(fiber.Map{
			"success": false,
			"message": "Forbidden",
		})
	}

	// Return if Unprocessable Entity
	if code == fiber.StatusUnprocessableEntity {
		return ctx.Status(code).JSON(fiber.Map{
			"success":    false,
			"message":    "Unprocessable Entity",
			"error_code": "UNPROCESSABLE_ENTITY",
			"error":      err.Error(),
		})
	}

	// Logging Error
	logger.Error(err)

	// Return Internal Server Error
	return ctx.Status(code).JSON(fiber.Map{
		"success": false,
		"message": "Internal Server Error",
	})
}
