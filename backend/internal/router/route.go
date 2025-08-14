package router

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	"backend/internal/handler"
)

func NewRouter() *gin.Engine {
    r := gin.Default()

	// CORSミドルウェア追加 TODO:パッケージ分ける
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"}, // ← Next.js側のURL TODO: 環境変数
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
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

    r.GET("/habit/list", handler.GetHabitList)
	r.GET("/habit/:id", handler.GetHabit)
	r.POST("/habit/register", handler.RegisterHabit)
	r.PUT("/habit/:id/update", handler.UpdateHabit)
	r.DELETE("/habit/:id/delete", handler.DeleteHabit)

    return r
}