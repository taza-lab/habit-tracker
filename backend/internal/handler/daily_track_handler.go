package handler

import (
	"net/http"

    "github.com/gin-gonic/gin"
	"backend/internal/domain/model/habit_track"
	"backend/internal/domain/model/habit"
	"backend/internal/domain/repository"
)

type DailyTrackHandler struct {
	habitTrackRepo repository.HabitTrackRepository
}

func NewDailyTrackHandler(repo repository.HabitTrackRepository) *DailyTrackHandler {
	return &DailyTrackHandler{
		habitTrackRepo: repo,
	}
}

func (h *DailyTrackHandler) GetDailyTrack(c *gin.Context) {
	var data = habit_track.DailyTrack{
		Date: c.Param("date"),
		Habits: []habit_track.HabitStatus{
			{Habit: habit.Habit{Id: 1, Name: "朝ヨガ"}, IsDone: false},
			{Habit: habit.Habit{Id: 2, Name: "勉強"}, IsDone: false},
		},
	}

	c.JSON(http.StatusOK, data)
}

func (h *DailyTrackHandler) UpdateDoneDailyTrack(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "success"})
}
