package handler

// handlerの規約
// フロントで表示するメッセージはここに定義
// リポジトリからのエラーはlog.Printf("[ERROR] ~")でそのまま出力

import (
	"errors"
	"log"
	"net/http"

	"backend/internal/domain/common"
	"backend/internal/domain/model/habit"
	"backend/internal/domain/repository"
	"backend/internal/utils"

	"github.com/gin-gonic/gin"
)

type HabitHandler struct {
	habitRepo repository.HabitRepository
}

func NewHabitHandler(repo repository.HabitRepository) *HabitHandler {
	return &HabitHandler{
		habitRepo: repo,
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
	habits, err := h.habitRepo.FetchAll(userId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "error"})
		return
	}

	if habits == nil {
		habits = make([]*habit.Habit, 0)
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

	// 新規登録
	newHabit := habit.Habit{UserId: userId, Name: habitRequest.Name}
	habit, err := h.habitRepo.Register(&newHabit)

	if err != nil {
		if errors.Is(err, common.ErrAlreadyExists) {
			c.JSON(http.StatusBadRequest, gin.H{"message": "すでに登録済みの習慣です。"})
			return
		}

		log.Printf("[ERROR] %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "エラーが発生しました。"})
		return
	}

	// TODO: 今日のdaily-trackに追加

	c.JSON(http.StatusOK, gin.H{"message": "success", "id": habit.Id})
}

func (h *HabitHandler) UpdateHabit(c *gin.Context) {
	// TODO: 今後使うから残しておく
	c.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (h *HabitHandler) DeleteHabit(c *gin.Context) {
	id := c.Param("id")

	// idが空文字列の場合のチェック
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "IDは必須です。"})
		return
	}

	// 削除
	err := h.habitRepo.Delete(id)

	if err != nil {
		log.Printf("[ERROR] %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "エラーが発生しました。"})
		return
	}

	// TODO: !isDoneだったら今日のdaily-trackから削除

	c.JSON(http.StatusOK, gin.H{"message": "success"})
}
