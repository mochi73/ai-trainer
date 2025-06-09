package domain

import "context"

type ChatService interface {
	GetCompletion(ctx context.Context, prompt string) (string, error)
}
