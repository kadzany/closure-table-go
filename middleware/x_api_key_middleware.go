package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kadzany/closure-table-go/config"
	"github.com/kadzany/closure-table-go/model/dto"
)

func XApiKeyMiddleware(ctx *fiber.Ctx) error {
	// Get Config
	env := config.GetEnvConfig()

	// Get Header
	if env.Get("X_API_KEY") != ctx.Get("X-API-Key") {
		return ctx.Status(fiber.StatusUnauthorized).JSON(dto.ApiResponseFail{
			Success: false,
			Message: "Unauthorized",
		})
	}

	// Next
	return ctx.Next()
}
