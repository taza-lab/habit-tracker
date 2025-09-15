package repositoryImpl

import (
	"context"
	"time"
	"fmt"

	"golang.org/x/crypto/bcrypt"
	
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	userModel "backend/internal/domain/model/user"
	"backend/internal/domain/repository"
	"backend/internal/domain/common"
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
		return nil, fmt.Errorf("failed to find user: %w", err)
	}

	// 取得したuserDBをドメインモデルのUserに変換
    user := &userModel.User{
        Id:       userDB.ID.Hex(), // ObjectIDをstringに変換
        Username: userDB.Username,
        Password: userDB.Password,
        Points:   userDB.Points,
    }

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
		return nil, fmt.Errorf("failed to find user by username: %w", err)
	}

	// 取得したuserDBをドメインモデルのUserに変換
    user := &userModel.User{
        Id:       userDB.ID.Hex(), // ObjectIDをstringに変換
        Username: userDB.Username,
        Password: userDB.Password,
        Points:   userDB.Points,
    }

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

	// ドメインモデルからDBモデルへ変換
    userDB := userDB{
        Username: user.Username,
        Password: string(hashedPassword),
        Points:   user.Points,
    }

	result, err := r.collection.InsertOne(ctx, userDB)
	if err != nil {
		return nil, fmt.Errorf("failed to register user: %w", err)
	}

	if oid, ok := result.InsertedID.(primitive.ObjectID); ok {
		user.Id = oid.Hex()
	}

	return user, nil
}

