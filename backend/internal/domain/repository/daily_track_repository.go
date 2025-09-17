package repository

import (
	"backend/internal/domain/model/daily_track"
)

type DailyTrackRepository interface {
	FindDailyTrack(userId string, targetDate string) (*daily_track.DailyTrack, error)
	RegisterDailyTrack(dailyTrack *daily_track.DailyTrack) (*daily_track.DailyTrack, error)
	UpdateHabitStatuses(dailyTrack *daily_track.DailyTrack) error
}
