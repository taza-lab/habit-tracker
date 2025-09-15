package repository

import (
	"backend/internal/domain/model/habit"
)

type HabitRepository interface {
	FetchAll() ([]habit.Habit, error)
}
