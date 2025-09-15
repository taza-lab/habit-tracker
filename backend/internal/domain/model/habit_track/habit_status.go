package habit_track

import (
    "backend/internal/domain/model/habit"
)

type HabitStatus struct {
    Habit  habit.Habit `json:"habit"`
    IsDone bool  `json:"isDone"`
}
