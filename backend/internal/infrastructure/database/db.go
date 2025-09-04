// infrastructure/database/db.go
package database

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// DBClient はMongoDBクライアントのインターフェース
type DBClient interface {
	Connect(ctx context.Context) error
	Disconnect(ctx context.Context) error
	Client() *mongo.Client
}

type mongoClient struct {
	client *mongo.Client
	uri    string
}

// NewDBClient は新しいMongoDBクライアントを作成
func NewDBClient(uri string) DBClient {
	return &mongoClient{
		uri: uri,
	}
}

// Connect はMongoDBに接続
func (m *mongoClient) Connect(ctx context.Context) error {
	client, err := mongo.NewClient(options.Client().ApplyURI(m.uri))
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	if err := client.Connect(ctx); err != nil {
		return err
	}
	// 接続確認
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		return err
	}

	m.client = client
	return nil
}

// Disconnect はDBから切断
func (m *mongoClient) Disconnect(ctx context.Context) error {
	if m.client == nil {
		return nil
	}
	return m.client.Disconnect(ctx)
}

// Client はmongo.Clientインスタンスを返す
func (m *mongoClient) Client() *mongo.Client {
	return m.client
}
