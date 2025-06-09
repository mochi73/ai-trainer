package domain

import "context"

type ChatUseCase interface {
	Prompt(ctx context.Context, message string) (string, error)
}
