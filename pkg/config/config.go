package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Printf("Ошибка загрузки .env файла: %v", err)
	}
}

func GetOpenAIAPIKey() string {
	return os.Getenv("OPENAI_API_KEY")
}
