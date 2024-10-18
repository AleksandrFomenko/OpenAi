package main

import (
	"OpenAi/pkg/config"
	"OpenAi/pkg/controllers"
	"OpenAi/pkg/middleware"
	"OpenAi/pkg/routes"
	"OpenAi/pkg/services"
	"OpenAi/pkg/services/OpenAI"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
	"github.com/gofiber/fiber/v3/middleware/logger"
)

func main() {

	config.LoadEnv()

	app := fiber.New()

	// Применяем middleware логирования ко всем маршрутам
	app.Use(logger.New(logger.Config{
		Format:     "${time} - ${ip} - ${method} ${path} - ${status} - ${latency}\n",
		TimeFormat: "02-Jan-2006",
		TimeZone:   "Local",
	}))

	app.Get("/", hello)

	apiGroup := app.Group("/api", middleware.APIKeyAuth())

	var aiService services.AiServices = &OpenAI.OpenAi{}
	openAIController := controllers.NewAiController(aiService)

	routes.RegisterOpenAIRoutes(apiGroup, openAIController)

	log.Fatal(app.Listen(":8080"))

}

func hello(c fiber.Ctx) error {
	return c.SendString("Hello, World!")
}
