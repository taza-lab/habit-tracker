package repository

import (
	"backend/internal/domain/model/habit"
	"context"
)

type HabitRepository interface {
	FetchAll(ctx context.Context, userId string) ([]*habit.Habit, error)
	Register(ctx context.Context, habit *habit.Habit) (*habit.Habit, error)
	Delete(ctx context.Context, id string) error
}
