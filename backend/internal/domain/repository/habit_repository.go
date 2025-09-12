package repository

import (
	"backend/internal/domain/model/habit"
)

type HabitRepository interface {
	FetchAll() ([]habit.Habit, error)
	Register(habit *habit.Habit) (*habit.Habit, error)
	Delete(id string) (error)
}
