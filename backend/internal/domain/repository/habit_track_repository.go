package repository

import (
	"backend/internal/domain/model/habit_track"
)

type HabitTrackRepository interface {
	FindDailyTrack(targetDate string) (habit_track.DailyTrack, error)
}
