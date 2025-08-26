package main

import (
    "log"
    "github.com/joho/godotenv"
    "backend/internal/router"
)

func init() {
    // .envの環境変数読み込み
    if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables.")
	}
}

func main() {
    r := router.NewRouter()
    r.Run(":8080")
}
