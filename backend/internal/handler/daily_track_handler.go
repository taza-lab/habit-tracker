package handler

import (
	"net/http"

    "github.com/gin-gonic/gin"
	"backend/internal/domain/model/habit_track"
)

func GetDailyTrack(c *gin.Context) {
	var data = habit_track.DailyTrack{
		Date: c.Param("date"),
		Habits: []habit_track.HabitStatus{
			{Habit: habit_track.Habit{Id: 1, Name: "朝ヨガ"}, IsDone: false},
			{Habit: habit_track.Habit{Id: 2, Name: "勉強"}, IsDone: false},
		},
	}

	c.JSON(http.StatusOK, data)
}

func UpdateDoneDailyTrack(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "success"})
}
