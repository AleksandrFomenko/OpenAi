package routes

import (
	"OpenAi/pkg/controllers"
	"github.com/gofiber/fiber/v3"
)

func RegisterOpenAIRoutes(api fiber.Router, controller *controllers.AiController) {
	api.Post("/openai", controller.TextHandler)
}
