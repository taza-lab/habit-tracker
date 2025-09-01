package handler

import (
	"net/http"
	"time"

    "github.com/gin-gonic/gin"
	"backend/internal/domain/habit_track"
)

// メモ
// gin.H = map[string]interface{}

func GetHabitList(c *gin.Context) {
	var data = []habit_track.Habit{
		{Id: 1, Name: "朝ヨガ"},
		{Id: 2, Name: "勉強"},
	}

	c.JSON(http.StatusOK, data)
}

func GetHabit(c *gin.Context) {
	var data = habit_track.Habit{Id: 1, Name: "朝ヨガ"}

	c.JSON(http.StatusOK, data)
}

func RegisterHabit(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "success", "id": time.Now().Format("20060102150405")})
}

func UpdateHabit(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "success"})
}

func DeleteHabit(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "success"})
}