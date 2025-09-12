package handler

import (
	"net/http"

    "github.com/gin-gonic/gin"
	"backend/internal/domain/model/daily_track"
	"backend/internal/domain/repository"
)

type DailyTrackHandler struct {
	dailyTrackRepo repository.DailyTrackRepository
}

func NewDailyTrackHandler(repo repository.DailyTrackRepository) *DailyTrackHandler {
	return &DailyTrackHandler{
		dailyTrackRepo: repo,
	}
}

func (h *DailyTrackHandler) GetDailyTrack(c *gin.Context) {
	var data = daily_track.DailyTrack{
		Date: c.Param("date"),
		HabitStatuses: []daily_track.HabitStatus{
			{HabitId: "1", IsDone: false},
			{HabitId: "2", IsDone: false},
		},
	}

	c.JSON(http.StatusOK, data)
}

func (h *DailyTrackHandler) UpdateDoneDailyTrack(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "success"})
}
