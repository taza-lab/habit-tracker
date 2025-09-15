package repositoryImpl

// RepositoryImpl規約
//  取得したドキュメントをドメインモデルに変換して返却する
//  エラーはDBからの元のエラーをfmt.Errorfでラップして返却する
//  想定外のエラーの場合はlogでログ出力 -> [ERROR] HabitRepository.FugaMethod ~

// メモ
// GoのPrintf系関数では、%v（値）、%+v（フィールド名付きの値）、%#v（Goの構文形式）といったフォーマット指定子を使うことで、構造体の内容をまとめて出力できます。

import (
	"context"
	"fmt"
	"log"
	"time"

	"backend/internal/domain/common"
	"backend/internal/domain/model/habit"
	"backend/internal/domain/repository"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// DBに保存するための内部モデル
type habitDB struct {
	ID     primitive.ObjectID `bson:"_id,omitempty"`
	UserId string             `bson:"user_id"`
	Name   string             `bson:"name"`
}

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

// 習慣一覧取得
func (r *HabitRepository) FetchAll(userId string) ([]habit.Habit, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Find()で全件取得
	cursor, err := r.collection.Find(ctx, bson.M{"user_id": userId})
	if err != nil {
		log.Printf("[ERROR] HabitRepository.FetchAll() failed to collection.Find (user_id: %s): %w", userId, err)

		// NOTE: nilスライスは要素が一つもない有効なスライスと認識される
		return nil, fmt.Errorf("failed to habit fetch all: %w", err)
	}
	defer cursor.Close(ctx)

	// 結果を格納するスライス
	var habits []habit.Habit
	if err = cursor.All(ctx, &habits); err != nil {
		log.Printf("[ERROR] HabitRepository.FetchAll() failed to cursor.All : %w", err)
		return nil, fmt.Errorf("failed to decode documents from cursor: %w", err)
	}

	return habits, nil
}

// 習慣登録
func (r *HabitRepository) Register(habit *habit.Habit) (*habit.Habit, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// nameの重複チェック
	var existingHabitDB habitDB
	err := r.collection.FindOne(ctx, bson.M{"user_id": habit.UserId, "name": habit.Name}).Decode(&existingHabitDB)

	if err == nil {
		// すでに同名のhabitが存在する場合はエラーを返す
		return nil, common.ErrAlreadyExists
	}
	if err != mongo.ErrNoDocuments {
		log.Printf("[ERROR] HabitRepository.Register() failed to collection.FindOne (name: %s) : %w", habit.Name, err)
		return nil, fmt.Errorf("failed to check for existing habit: %w", err)
	}

	// DBに保存するためのモデルに変換
	habitDB := habitDB{
		UserId: habit.UserId,
		Name:   habit.Name,
	}

	// 新規登録
	result, err := r.collection.InsertOne(ctx, habitDB)

	if err != nil {
		log.Printf("[ERROR] HabitRepository.Register() failed to collection.InsertOne (data: %+v) : %w", habitDB, err)
		return nil, fmt.Errorf("failed to register habit: %w", err)
	}

	if oid, ok := result.InsertedID.(primitive.ObjectID); ok {
		habit.Id = oid.Hex()
	}

	return habit, nil
}

// 習慣削除
func (r *HabitRepository) Delete(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var result *mongo.DeleteResult
	var err error

	// MongoDBの_idはObjectID型で保存される
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Printf("[ERROR] HabitRepository.Delete() failed to primitive.ObjectIDFromHex (id: %s) : %w", id, err)
		return fmt.Errorf("invalid ID: %w", err)
	}

	result, err = r.collection.DeleteOne(ctx, bson.M{"_id": objectID})
	if err != nil {
		log.Printf("[ERROR] HabitRepository.Delete() failed to collection.DeleteOne (_id: %s) : %w", id, err)
		return fmt.Errorf("failed to delete habit: %w", err)
	}

	if result.DeletedCount == 0 {
		return common.ErrNotFound
	}

	return nil
}
