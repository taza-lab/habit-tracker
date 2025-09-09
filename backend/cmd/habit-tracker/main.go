package main

import (
    "context"
	"log"
	"os"
	"time"

    "github.com/joho/godotenv"
    "backend/internal/infrastructure/database"
    "backend/internal/router"
)

func init() {
    // .envの環境変数読み込み
    if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables.")
	}
}

func main() {
    // DB接続

    // 依存性の解決
	// 1. 環境変数からDB URIを取得
    dbUri := os.Getenv("DATABASE_URI")

    // 2. DBクライアントを作成
	dbClient := database.NewDBClient(dbUri)

    // 3. DBに接続
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := dbClient.Connect(ctx); err != nil {
		log.Fatal("Could not connect to DB:", err)
	}
	defer dbClient.Disconnect(ctx)
	log.Println("Connected to DB!")

    // Route
    r := router.NewRouter()
    r.Run(":8080")
}
