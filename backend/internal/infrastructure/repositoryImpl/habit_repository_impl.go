package repositoryImpl

import (
	"context"
	"time"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"backend/internal/domain/model/habit"
	"backend/internal/domain/repository"
)

// HabitRepository はMongoDBのusersコレクションにアクセスします
type HabitRepository struct {
	collection *mongo.Collection
}

// NewHabitRepository は新しいHabitRepositoryインスタンスを作成します
func NewHabitRepository(collection *mongo.Collection) repository.HabitRepository {
	return &HabitRepository{
		collection: collection, // habit collection
	}
}

func (r *HabitRepository) FetchAll() ([]habit.Habit, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Find()で全件取得
	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		// NOTE: nilスライスは要素が一つもない有効なスライスと認識される
		return nil, fmt.Errorf("failed to habit fetch all: %w", err)
	}
	defer cursor.Close(ctx)

	// 結果を格納するスライス
	var habits []habit.Habit
	if err = cursor.All(ctx, &habits); err != nil {
		return nil, fmt.Errorf("failed to decode documents from cursor: %w", err)
	}

	return habits, nil
}
