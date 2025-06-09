package openai

import (
	"ai_trainer/internal/domain"
	"context"
	"log"
	"os"

	openai "github.com/sashabaranov/go-openai"
)

type openAIClient struct {
	Client *openai.Client
}

func NewOpenAIClient() domain.ChatService {
	apiKey := os.Getenv("OPENAI_API_KEY")

	return &openAIClient{
		Client: openai.NewClient(apiKey),
	}
}

func (c *openAIClient) GetCompletion(ctx context.Context, prompt string) (string, error) {
	resp, err := c.Client.CreateChatCompletion(ctx, openai.ChatCompletionRequest{
		Model: openai.GPT3Dot5Turbo,
		Messages: []openai.ChatCompletionMessage{
			{Role: "user", Content: prompt},
		},
	})

	if err != nil {
		log.Printf("OpenAIへのリクエスト失敗: %+v", err)
		return "", err
	}

	return resp.Choices[0].Message.Content, nil

}
