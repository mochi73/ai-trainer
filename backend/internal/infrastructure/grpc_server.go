package infrastructure

import (
	trainerpf "ai_trainer/gen/trainerpf/v1"
	evaluateworkoutgrpc "ai_trainer/internal/adapter/evaluate_workout_grpc"
	"ai_trainer/internal/adapter/openai"
	"ai_trainer/internal/usecase"

	"google.golang.org/grpc"
)

func RegisterGRPCServices(s *grpc.Server) {
	client := openai.NewOpenAIClient()
	usecase := usecase.NewChatUseCase(client)
	handler := evaluateworkoutgrpc.NewChatHandler(usecase)
	trainerpf.RegisterWorkoutServiceServer(s, handler)
}
