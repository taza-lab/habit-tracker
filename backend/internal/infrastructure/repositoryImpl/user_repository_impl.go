package repositoryImpl

import (
	"context"
	"fmt"
	"log"
	"time"

	"golang.org/x/crypto/bcrypt"

	"backend/internal/domain/common"
	userModel "backend/internal/domain/model/user"
	"backend/internal/domain/repository"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// DBに保存するための内部モデル
type userDB struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Username string             `bson:"username"`
	Password string             `bson:"password"`
	Points   int                `bson:"points"`
}

// UserRepository はMongoDBのusersコレクションにアクセスします
type UserRepository struct {
	collection *mongo.Collection
}

// NewUserRepository は新しいUserRepositoryインスタンスを作成します
func NewUserRepository(collection *mongo.Collection) repository.UserRepository {
	return &UserRepository{
		collection: collection,
	}
}

func (r *UserRepository) Find(id string) (*userModel.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// MongoDBの_idはObjectID型で保存される
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("invalid ID: %w", err)
	}

	var userDB userDB
	err = r.collection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&userDB)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, common.ErrNotFound
		}

		log.Printf("[ERROR] UserRepository.Find() failed to collection.FindOne (user_id: %s): %w", id, err)
		return nil, fmt.Errorf("failed to find user: %w", err)
	}

	user := convertToUser(&userDB)

	return user, nil
}

func (r *UserRepository) FindByUserName(username string) (*userModel.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var userDB userDB
	err := r.collection.FindOne(ctx, bson.M{"username": username}).Decode(&userDB)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, common.ErrNotFound
		}
		log.Printf("[ERROR] UserRepository.Find() failed to collection.FindOne (usernaem: %s): %w", username, err)
		return nil, fmt.Errorf("failed to find user by username: %w", err)
	}

	user := convertToUser(&userDB)

	return user, nil
}

func (r *UserRepository) Register(user *userModel.User) (*userModel.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// パスワードハッシュ化
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %w", err)
	}

	// 登録用のBDモデルを作成
	userDB := userDB{
		Username: user.Username,
		Password: string(hashedPassword),
		Points:   user.Points,
	}

	result, err := r.collection.InsertOne(ctx, userDB)
	if err != nil {
		log.Printf("[ERROR] UserRepository.Register() failed to collection.InsertOne (data: %+v): %w", userDB, err)
		return nil, fmt.Errorf("failed to register user: %w", err)
	}

	if oid, ok := result.InsertedID.(primitive.ObjectID); ok {
		user.Id = oid.Hex()
	}

	return user, nil
}

func (r *UserRepository) UpdatePoints(userId string, points int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// ID変換
	objectID, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		log.Printf("[ERROR] UserRepository.UpdatePoints() failed to primitive.ObjectIDFromHex (id: %s) : %w", userId, err)
		return fmt.Errorf("invalid ID: %w", err)
	}

	// 更新対象を特定するフィルタ
	filter := bson.M{"_id": objectID}

	// 更新内容
	update := bson.M{
		"$set": bson.M{
			// ポイントフィールドのみ更新可能
			"points": points,
		},
	}

	var result *mongo.UpdateResult
	result, err = r.collection.UpdateOne(ctx, filter, update)

	if err != nil {
		log.Printf("[ERROR] UserRepository.UpdatePoints() failed to collection.UpdateOne (_id: %s, points: %s) : %w", userId, points, err)
		return fmt.Errorf("failed to update points: %w", err)
	}

	if result.MatchedCount == 0 {
		log.Printf("[ERROR] UserRepository.UpdatePoints() failed to collection.UpdateOne target not found (_id: %s)", userId)
		return common.ErrNotFound
	}

	return nil
}

// DBモデルをドメインモデルに変換
func convertToUser(userDB *userDB) *userModel.User {
	return &userModel.User{
		Id:       userDB.ID.Hex(), // ObjectIDをstringに変換
		Username: userDB.Username,
		Password: userDB.Password,
		Points:   userDB.Points,
	}
}
