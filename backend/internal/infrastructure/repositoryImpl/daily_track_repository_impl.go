package repositoryImpl

import (
	"context"
	"fmt"
	"log"
	"time"

	"backend/internal/domain/common"
	"backend/internal/domain/model/daily_track"
	"backend/internal/domain/repository"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// DBに保存するための内部モデル
type habitStatusDB struct {
	HabitId   string `bson:"habit_id"`
	HabitName string `bson:"habit_name"`
	IsDone    bool   `bson:"is_done"`
}
type dailyTrackDB struct {
	ID            primitive.ObjectID `bson:"_id,omitempty"`
	UserId        string             `bson:"user_id"`
	Date          string             `bson:"date"`
	HabitStatuses []habitStatusDB    `bson:"habit_statuses"`
}

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

func (r *DailyTrackRepository) FindDailyTrack(userId string, targetDate string) (*daily_track.DailyTrack, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var dailyTrackDB dailyTrackDB
	err := r.collection.FindOne(ctx, bson.M{"user_id": userId, "date": targetDate}).Decode(&dailyTrackDB)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, common.ErrNotFound
		}
		log.Printf("[ERROR] DailyTrackRepository.FindDailyTrack() failed to collection.FindOne (user_id: %s, date: %s): %w", userId, targetDate, err)
		return nil, fmt.Errorf("failed to find daily_track: %w", err)
	}

	dailyTrack := convertToDailyTrack(&dailyTrackDB)

	return dailyTrack, nil
}

func (r *DailyTrackRepository) RegisterDailyTrack(dailyTrack *daily_track.DailyTrack) (*daily_track.DailyTrack, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 登録済みチェック
	var existsDailyTrackDB dailyTrackDB
	err := r.collection.FindOne(ctx, bson.M{"user_id": dailyTrack.UserId, "date": dailyTrack.Date}).Decode(&existsDailyTrackDB)
	if err == nil {
		return nil, common.ErrAlreadyExists
	}
	if err != mongo.ErrNoDocuments {
		log.Printf("[ERROR] DailyTrackRepository.RegisterDailyTrack() failed to collection.FindOne (user_id: %s, date: %s): %w", dailyTrack.UserId, dailyTrack.Date, err)
		return nil, fmt.Errorf("failed to find daily_track: %w", err)
	}

	dailyTrackDB := convertToDailyTrackDBWithoutId(dailyTrack)
	result, err := r.collection.InsertOne(ctx, dailyTrackDB)

	if err != nil {
		log.Printf("[ERROR] DailyTrackRepository.RegisterDailyTrack() failed to collection.InsertOne (data: %+v) : %w", dailyTrackDB, err)
		return nil, fmt.Errorf("failed to register daily_track: %w", err)
	}

	if oid, ok := result.InsertedID.(primitive.ObjectID); ok {
		dailyTrack.Id = oid.Hex()
	}

	return dailyTrack, nil
}

func (r *DailyTrackRepository) UpdateHabitStatuses(dailyTrack *daily_track.DailyTrack) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// ID変換
	objectID, err := primitive.ObjectIDFromHex(dailyTrack.Id)
	if err != nil {
		log.Printf("[ERROR] DailyTrackRepository.UpdateHabitStatuses() failed to primitive.ObjectIDFromHex (id: %s) : %w", dailyTrack.Id, err)
		return fmt.Errorf("invalid ID: %w", err)
	}

	// DBモデルに変換
	dailyTrackDB := convertToDailyTrackDBWithoutId(dailyTrack)

	// 更新対象を特定するフィルタ
	filter := bson.M{"_id": objectID}

	// 更新内容
	update := bson.M{
		"$set": bson.M{
			// ステータスフィールドのみ更新可能
			"habit_statuses": dailyTrackDB.HabitStatuses,
		},
	}

	var result *mongo.UpdateResult
	result, err = r.collection.UpdateOne(ctx, filter, update)

	if err != nil {
		log.Printf("[ERROR] DailyTrackRepository.UpdateHabitStatuses() failed to collection.UpdateOne (_id: %s, statuses: %+v) : %w", dailyTrack.Id, dailyTrackDB.HabitStatuses, err)
		return fmt.Errorf("failed to update daily_track: %w", err)
	}

	if result.MatchedCount == 0 {
		log.Printf("[ERROR] DailyTrackRepository.UpdateHabitStatuses() failed to collection.UpdateOne target not found (_id: %s)", dailyTrack.Id)
		return common.ErrNotFound
	}

	return nil
}

// DBモデルをドメインモデルに変換
func convertToDailyTrack(dailyTrackDB *dailyTrackDB) *daily_track.DailyTrack {
	var habitStatuses []*daily_track.HabitStatus
	for _, habitStatusDB := range dailyTrackDB.HabitStatuses {
		habitStatus := &daily_track.HabitStatus{
			HabitId:   habitStatusDB.HabitId,
			HabitName: habitStatusDB.HabitName,
			IsDone:    habitStatusDB.IsDone,
		}
		habitStatuses = append(habitStatuses, habitStatus)
	}

	return &daily_track.DailyTrack{
		Id:            dailyTrackDB.ID.Hex(),
		UserId:        dailyTrackDB.UserId,
		Date:          dailyTrackDB.Date,
		HabitStatuses: habitStatuses,
	}
}

// ドメインモデルをDBモデルに変換
// NOTE: ID変換の責任は負わない
func convertToDailyTrackDBWithoutId(dailyTrack *daily_track.DailyTrack) *dailyTrackDB {
	var habitStatusesDB []habitStatusDB
	for _, habitStatus := range dailyTrack.HabitStatuses {
		habitStatus := habitStatusDB{
			HabitId:   habitStatus.HabitId,
			HabitName: habitStatus.HabitName,
			IsDone:    habitStatus.IsDone,
		}
		habitStatusesDB = append(habitStatusesDB, habitStatus)
	}

	return &dailyTrackDB{
		UserId:        dailyTrack.UserId,
		Date:          dailyTrack.Date,
		HabitStatuses: habitStatusesDB,
	}
}
