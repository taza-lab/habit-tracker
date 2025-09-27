package service

import (
	"backend/internal/domain/model/habit"
	"context"
)

type HabitService interface {
	GetHabitList(ctx context.Context, userId string) ([]*habit.Habit, error)
	RegisterHabit(ctx context.Context, userId string, habitName string) (*habit.Habit, error)
	DeleteHabit(ctx context.Context, userId string, habitId string) error
}
