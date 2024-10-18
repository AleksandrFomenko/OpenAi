package middleware

import (
	"github.com/gofiber/fiber/v3"
	"os"
)

func APIKeyAuth() fiber.Handler {
	clientAPIKey := os.Getenv("CLIENT_API_KEY")
	if clientAPIKey == "" {
		panic("CLIENT_API_KEY не задан в переменных окружения")
	}

	return func(c fiber.Ctx) error {
		apiKey := c.Get("Authorization")
		if apiKey == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "API-ключ не предоставлен",
			})
		}

		if apiKey != clientAPIKey {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Недействительный API-ключ",
			})
		}

		return c.Next()
	}
}
