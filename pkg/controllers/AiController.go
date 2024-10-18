package controllers

import (
	"OpenAi/pkg/services"
	"github.com/gofiber/fiber/v3"
	"log"
)

type AiController struct {
	AiServices services.AiServices
}

func NewAiController(aiService services.AiServices) *AiController {
	return &AiController{AiServices: aiService}
}

func (ctrl *AiController) TextHandler(c fiber.Ctx) error {

	var request struct {
		Prompt string `json:"prompt"`
	}

	if err := c.Bind().JSON(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Неверный формат запроса",
		})
	}

	log.Printf("Получен запрос: %s", request.Prompt)

	response, err := ctrl.AiServices.GetTextResponse(request.Prompt)
	if err != nil {
		log.Printf("Ошибка AI-сервиса: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	log.Printf("Ответ AI-сервиса: %s", response)

	return c.JSON(fiber.Map{
		"response": response,
	})
}
