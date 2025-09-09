package repositoryImpl

import (
	"context"
	"time"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"backend/internal/domain/model/habit_track"
	"backend/internal/domain/repository"
	"backend/internal/domain/common"
)

// HabitTrackRepository はMongoDBのusersコレクションにアクセスします
type HabitTrackRepository struct {
	collection *mongo.Collection
}

// NewHabitTrackRepository は新しいHabitTrackRepositoryインスタンスを作成します
func NewHabitTrackRepository(collection *mongo.Collection) repository.HabitTrackRepository {
	return &HabitTrackRepository{
		collection: collection,
	}
}

func (r *HabitTrackRepository) FindDailyTrack(targetDate string) (habit_track.DailyTrack, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var dailyTrack habit_track.DailyTrack
	err := r.collection.FindOne(ctx, bson.M{"date": targetDate}).Decode(&dailyTrack)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return habit_track.DailyTrack{}, common.ErrNotFound
		}
		return habit_track.DailyTrack{}, fmt.Errorf("failed to find user by name: %w", err)
	}

	return dailyTrack, nil
}
