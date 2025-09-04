package mongodb

import (
	"context"
	"your_project_name/domain/repository" // ドメイン層のインターフェースをインポート

	"go.mongodb.org/mongo-driver/mongo"
)

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

// FindAll は全てのユーザーを取得します
func (r *UserRepository) FindAll() ([]repository.User, error) {
	// ... MongoDBのクエリをここに書く
	// このメソッドはr.collectionを使って操作を行います
	return nil, nil
}
