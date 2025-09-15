package repositoryImpl

import (
	"context"
	"time"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"backend/internal/domain/model/daily_track"
	"backend/internal/domain/repository"
	"backend/internal/domain/common"
)

// DailyTrackRepository はMongoDBのusersコレクションにアクセスします
type DailyTrackRepository struct {
	collection *mongo.Collection
}

// NewDailyTrackRepository は新しいDailyTrackRepositoryインスタンスを作成します
func NewDailyTrackRepository(collection *mongo.Collection) repository.DailyTrackRepository {
	return &DailyTrackRepository{
		collection: collection,
	}
}

func (r *DailyTrackRepository) FindDailyTrack(targetDate string) (daily_track.DailyTrack, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var dailyTrack daily_track.DailyTrack
	err := r.collection.FindOne(ctx, bson.M{"date": targetDate}).Decode(&dailyTrack)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return daily_track.DailyTrack{}, common.ErrNotFound
		}
		return daily_track.DailyTrack{}, fmt.Errorf("failed to find user by name: %w", err)
	}

	return dailyTrack, nil
}
