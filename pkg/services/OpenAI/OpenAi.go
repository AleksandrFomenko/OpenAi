package OpenAI

import (
	"OpenAi/pkg/config"
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

type OpenAi struct {
}

func (o *OpenAi) GetTextResponse(prompt string) (string, error) {
	apiKey := config.GetOpenAIAPIKey()
	if apiKey == "" {
		return "", errors.New("API-ключ OpenAI не задан")
	}

	url := "https://api.openai.com/v1/chat/completions"

	requestBody, err := json.Marshal(map[string]interface{}{
		"model": "gpt-4o",
		"messages": []map[string]string{
			{"role": "user", "content": prompt},
		},
	})
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	if resp.StatusCode != http.StatusOK {
		return "", errors.New("Ошибка от OpenAI API: " + resp.Status)
	}

	var response struct {
		Choices []struct {
			Message struct {
				Content string `json:"content"`
			} `json:"message"`
		} `json:"choices"`
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		return "", err
	}

	if len(response.Choices) == 0 {
		return "", errors.New("Пустой ответ от OpenAI API")
	}

	return response.Choices[0].Message.Content, nil
}
