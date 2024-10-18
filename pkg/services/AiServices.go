package services

type AiServices interface {
	GetTextResponse(promt string) (string, error)
}
