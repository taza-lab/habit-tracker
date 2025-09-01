package router

import (
	"net/http"
	"time"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	"backend/internal/handler"
	"backend/internal/middleware"
)

func NewRouter() *gin.Engine {
    r := gin.Default()

	// CORSミドルウェア追加 TODO:パッケージ分ける
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{os.Getenv("NEXT_BASE_URL")},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// ヘルスチェック
	r.GET("/health", func (c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
				"message": "OK",

		})
	})

	// ログイン
	r.POST("/login", handler.Login)

	protected := r.Group("/auth")
	protected.Use(middleware.AuthMiddleware())
	{
		// ユーザー
		protected.GET("/user", handler.GetUser) 

		// 習慣トラック
		protected.GET("/daily_track/:date", handler.GetDailyTrack)
		protected.POST("/daily_track/done", handler.UpdateDoneDailyTrack)

		// 習慣の管理
		protected.GET("/habit/list", handler.GetHabitList)
		protected.GET("/habit/:id", handler.GetHabit)
		protected.POST("/habit/register", handler.RegisterHabit)
		protected.PUT("/habit/:id/update", handler.UpdateHabit)
		protected.DELETE("/habit/:id/delete", handler.DeleteHabit)
	}

    return r
}