package main

import (
	"context"
	"log"
	"os"
	"time"

	"backend/internal/handler"
	"backend/internal/infrastructure/database"
	"backend/internal/infrastructure/repositoryImpl"
	"backend/internal/infrastructure/serviceImpl"
	"backend/internal/middleware"
	"backend/internal/router"

	"github.com/joho/godotenv"
)

func init() {
	// .envの環境変数読み込み
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables.")
	}

	// 環境変数の初期化を安全に行う
	middleware.InitJWTSecret()
}

func main() {
	// --- DB接続 ---
	// 1. 環境変数からDB URIを取得
	dbUri := os.Getenv("DATABASE_URI")
	dbName := os.Getenv("DATABASE_NAME")
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

	// --- 依存性の解決とインスタンス化 ---
	// 1. 各リポジトリを生成し、DBクライアントを注入
	db := dbClient.Client().Database(dbName)
	userRepo := repositoryImpl.NewUserRepository(db.Collection("user"))
	habitRepo := repositoryImpl.NewHabitRepository(db.Collection("habits"))
	dailyTrackRepo := repositoryImpl.NewDailyTrackRepository(db.Collection("daily_track"))

	// 2. 各サービスを生成し、使用するリポジトリを注入
	userService := serviceImpl.NewUserService(dbClient.Client(), userRepo)
	habitService := serviceImpl.NewHabitService(dbClient.Client(), habitRepo, dailyTrackRepo)
	dailyTrackService := serviceImpl.NewDailyTrackService(dbClient.Client(), userRepo, habitRepo, dailyTrackRepo)

	// 3. 各ハンドラーを生成し、対応するサービスを注入
	userHandler := handler.NewUserHandler(userService)
	habitHandler := handler.NewHabitHandler(habitService)
	dailyTrackHandler := handler.NewDailyTrackHandler(dailyTrackService)

	// 4. ルーター設定のコンフィグを作成
	routerConfig := &router.RouterConfig{
		UserHandler:       userHandler,
		HabitHandler:      habitHandler,
		DailyTrackHandler: dailyTrackHandler,
	}

	// Route
	r := router.NewRouter(routerConfig)
	r.Run(":8080")
}
