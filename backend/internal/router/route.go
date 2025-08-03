package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"backend/internal/handler"
)

func NewRouter() *gin.Engine {
    r := gin.Default()

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