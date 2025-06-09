package evaluateworkoutgrpc

import (
	trainerpf "ai_trainer/gen/trainerpf/v1"
	"ai_trainer/internal/domain/model"
	"ai_trainer/internal/usecase"
	"context"
)

type ChatHandler struct {
	trainerpf.UnimplementedWorkoutServiceServer
	uc usecase.ChatUseCase
}

func NewChatHandler(uc usecase.ChatUseCase) *ChatHandler {
	return &ChatHandler{uc: uc}
}

func (h *ChatHandler) EvaluateWorkout(ctx context.Context, req *trainerpf.EvaluateWorkoutRequest) (*trainerpf.EvaluateWorkoutResponse, error) {
	workouts := make([]*model.Workout, 0, len(req.Workouts))
	for _, w := range req.Workouts {
		sets := make([]*model.Set, 0, len(w.Sets))
		for _, s := range w.Sets {
			sets = append(sets, &model.Set{
				Load: s.Load,
				Reps: s.Reps,
			})
		}
		workouts = append(workouts, &model.Workout{
			Name: w.Name,
			Sets: sets,
		})
	}
	ew := model.EvaluateWorkout{
		Age:      req.Age,
		Weight:   req.Weight,
		Workouts: workouts,
	}
	advice, err := h.uc.EvaluateWorkout(ctx, ew)
	if err != nil {
		return nil, err
	}
	return &trainerpf.EvaluateWorkoutResponse{Advice: advice}, nil
}
