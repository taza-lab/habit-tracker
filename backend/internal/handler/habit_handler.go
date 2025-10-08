package handler

// handler規約
// フロントで表示するメッセージはここに定義
// サービスからのエラーはlog.Printf("[ERROR] ~")でそのまま出力

import (
	"errors"
	"log"
	"net/http"

	"backend/internal/domain/common"
	"backend/internal/domain/service"
	"backend/internal/utils"

	"github.com/gin-gonic/gin"
)

type HabitHandler struct {
	habitService service.HabitService
}

func NewHabitHandler(habitService service.HabitService) *HabitHandler {
	return &HabitHandler{
		habitService: habitService,
	}
}

// TODO: requestパッケージ作成
type HabitRequest struct {
	Id   string `json:"id"   binding:"required"`
	Name string `json:"name" binding:"required"`
}

// メモ
// gin.H = map[string]interface{}

func (h *HabitHandler) GetHabitList(c *gin.Context) {

	userId := utils.GetUserIdFromContext(c)
	habits, err := h.habitService.GetHabitList(c.Request.Context(), userId)

	if err != nil {
		log.Printf("[ERROR] %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "エラーが発生しました。"})
		return
	}

	c.JSON(http.StatusOK, habits)
}

func (h *HabitHandler) RegisterHabit(c *gin.Context) {
	userId := utils.GetUserIdFromContext(c)

	// バリデーション
	var habitRequest HabitRequest
	if err := c.ShouldBindJSON(&habitRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request body"})
		return
	}

	habit, err := h.habitService.RegisterHabit(c.Request.Context(), userId, habitRequest.Name)

	if err != nil {
		if errors.Is(err, common.ErrAlreadyExists) {
			c.JSON(http.StatusBadRequest, gin.H{"message": "すでに登録済みの習慣です。"})
			return
		}

		log.Printf("[ERROR] %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "エラーが発生しました。"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "id": habit.Id})
}

func (h *HabitHandler) DeleteHabit(c *gin.Context) {
	userId := utils.GetUserIdFromContext(c)
	targetHabitId := c.Param("id")

	// idが空文字列の場合のチェック
	if targetHabitId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "IDは必須です。"})
		return
	}
	// 削除
	err := h.habitService.DeleteHabit(c.Request.Context(), userId, targetHabitId)

	if err != nil {
		log.Printf("[ERROR] %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "エラーが発生しました。"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success"})
}
