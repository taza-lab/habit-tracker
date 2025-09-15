package router

import (
	"net/http"
	
	"github.com/gin-gonic/gin"
	"backend/internal/handler"
	"backend/internal/middleware"
)

// 必要な依存性をまとめた構造体
type RouterConfig struct {
    UserHandler *handler.UserHandler
    HabitHandler *handler.HabitHandler
    DailyTrackHandler *handler.DailyTrackHandler
}

func NewRouter(config *RouterConfig) *gin.Engine {
    r := gin.Default()

	r.Use(middleware.CorsMiddleware())

	// ヘルスチェック
	r.GET("/health", func (c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
				"message": "OK",

		})
	})

	// サインアップ
	r.POST("/signup", config.UserHandler.SignUp)

	// ログイン
	r.POST("/login", config.UserHandler.Login)

	protected := r.Group("/auth")
	protected.Use(middleware.AuthMiddleware())
	{
		// ユーザー
		protected.GET("/user", config.UserHandler.GetUser) 

		// 習慣トラック
		protected.GET("/daily_track/:date", config.DailyTrackHandler.GetDailyTrack)
		protected.POST("/daily_track/done", config.DailyTrackHandler.UpdateDoneDailyTrack)

		// 習慣の管理
		protected.GET("/habit/list", config.HabitHandler.GetHabitList)
		protected.POST("/habit/register", config.HabitHandler.RegisterHabit)
		protected.PUT("/habit/:id/update", config.HabitHandler.UpdateHabit)
		protected.DELETE("/habit/:id/delete", config.HabitHandler.DeleteHabit)
	}

    return r
}
