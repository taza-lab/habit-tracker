package repository

import (
	"backend/internal/domain/model/daily_track"
	"context"
)

type DailyTrackRepository interface {
	FindDailyTrack(ctx context.Context, userId string, targetDate string) (*daily_track.DailyTrack, error)
	RegisterDailyTrack(ctx context.Context, dailyTrack *daily_track.DailyTrack) (*daily_track.DailyTrack, error)
	UpdateHabitStatuses(ctx context.Context, dailyTrack *daily_track.DailyTrack) error
}
