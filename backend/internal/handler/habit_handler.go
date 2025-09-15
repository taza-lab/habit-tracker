package handler

import (
	"net/http"
	"time"

    "github.com/gin-gonic/gin"
	"backend/internal/domain/model/habit"
	"backend/internal/domain/repository"
)

type HabitHandler struct {
	habitRepo repository.HabitRepository
}

func NewHabitHandler(repo repository.HabitRepository) *HabitHandler {
	return &HabitHandler{
		habitRepo: repo,
	}
}

// メモ
// gin.H = map[string]interface{}

func (h *HabitHandler) GetHabitList(c *gin.Context) {
	habits, err := h.habitRepo.FetchAll()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "error"})
		return
	}

	// FetchAll() が正常に完了したが、結果がnilだった場合
	// クライアントに空のJSON配列を返す
	if habits == nil {
		habits = make([]habit.Habit, 0)
	}

	c.JSON(http.StatusOK, habits)
}

func (h *HabitHandler) GetHabit(c *gin.Context) {
	var data = habit.Habit{Id: 1, Name: "朝ヨガ"}

	c.JSON(http.StatusOK, data)
}

func (h *HabitHandler) RegisterHabit(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "success", "id": time.Now().Format("20060102150405")})
}

func (h *HabitHandler) UpdateHabit(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (h *HabitHandler) DeleteHabit(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "success"})
}
