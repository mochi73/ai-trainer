package usecase

import (
	"ai_trainer/internal/domain"
	"ai_trainer/internal/domain/model"
	"bytes"
	"context"
	"html/template"
)

type ChatUseCase interface {
	EvaluateWorkout(ctx context.Context, request model.EvaluateWorkout) (string, error)
}

type chatUseCase struct {
	ChatService domain.ChatService
}

func NewChatUseCase(service domain.ChatService) ChatUseCase {
	return &chatUseCase{
		ChatService: service,
	}
}

func (u *chatUseCase) EvaluateWorkout(ctx context.Context, request model.EvaluateWorkout) (string, error) {
	msg, err := BuildPrompt(request)
	if err != nil {
		return "", err
	}
	return u.ChatService.GetCompletion(ctx, msg)
}

func BuildPrompt(input model.EvaluateWorkout) (string, error) {
	tmpl, err := template.New("prompt").Parse(model.PromptTemplate)
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, input); err != nil {
		return "", err
	}

	return buf.String(), nil
}
