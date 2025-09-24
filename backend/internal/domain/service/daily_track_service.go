package service

import (
	"backend/internal/domain/model/daily_track"
	"context"
)

type DailyTrackService interface {
	GetDailyTrack(ctx context.Context, userId string, targetDate string) (*daily_track.DailyTrack, error)
	UpdateDoneDailyTrack(ctx context.Context, userId string, targetDate string, targetHabitId string) error
}
