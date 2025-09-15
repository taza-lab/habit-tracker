package repository

import (
	"backend/internal/domain/model/daily_track"
)

type DailyTrackRepository interface {
	FindDailyTrack(targetDate string) (daily_track.DailyTrack, error)
}
