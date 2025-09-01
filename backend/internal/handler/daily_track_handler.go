package handler

import (
	"net/http"

    "github.com/gin-gonic/gin"
	"backend/internal/domain/model"
)

func GetDailyTrack(c *gin.Context) {
	var data = model.DailyTrack{
		Date: c.Param("date"),
		Habits: []model.HabitStatus{
			{Habit: model.Habit{Id: 1, Name: "朝ヨガ"}, IsDone: false},
			{Habit: model.Habit{Id: 2, Name: "勉強"}, IsDone: false},
		},
	}

	c.JSON(http.StatusOK, data)
}

func UpdateDoneDailyTrack(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "success"})
}