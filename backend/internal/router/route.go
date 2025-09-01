package router

import (
	"net/http"
	
	"github.com/gin-gonic/gin"
	"backend/internal/handler"
	"backend/internal/middleware"
)

func NewRouter() *gin.Engine {
    r := gin.Default()

	r.Use(middleware.CorsMiddleware())

	// ヘルスチェック
	r.GET("/health", func (c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
				"message": "OK",

		})
	})

	// サインアップ
	r.POST("/signup", handler.SignUp)

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