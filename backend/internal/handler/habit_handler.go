package handler

// handlerの規約
// フロントで表示するメッセージはここに定義
// リポジトリからのエラーはlog.Printf("[ERROR] ~")でそのまま出力

import (
	"errors"
	"log"
	"net/http"
	"time"

	"backend/internal/domain/common"
	"backend/internal/domain/model/daily_track"
	"backend/internal/domain/model/habit"
	"backend/internal/domain/repository"
	"backend/internal/utils"

	"github.com/gin-gonic/gin"
)

type HabitHandler struct {
	habitRepo      repository.HabitRepository
	dailyTrackRepo repository.DailyTrackRepository
}

func NewHabitHandler(habitRepository repository.HabitRepository, dailyTrackRepository repository.DailyTrackRepository) *HabitHandler {
	return &HabitHandler{
		habitRepo:      habitRepository,
		dailyTrackRepo: dailyTrackRepository,
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

	// 今日のdaily-trackを取得
	todayString := time.Now().Format(`2006-01-02`) // YYYY-MM-DD
	log.Printf("[DEBUG] %v", todayString)
	todaysTrack, err := h.dailyTrackRepo.FindDailyTrack(userId, todayString)
	if err != nil && !errors.Is(err, common.ErrNotFound) {
		log.Printf("[ERROR] %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "エラーが発生しました。"})
		return
	}

	// 今日のdaily-trackがあれば作成した習慣を追加
	if !errors.Is(err, common.ErrNotFound) {

		newHabitStatus := &daily_track.HabitStatus{
			HabitId:   newHabit.Id,
			HabitName: newHabit.Name,
			IsDone:    false,
		}
		todaysTrack.HabitStatuses = append(todaysTrack.HabitStatuses, newHabitStatus)

		err = h.dailyTrackRepo.UpdateHabitStatuses(todaysTrack)
		if err != nil {
			log.Printf("[ERROR] %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"message": "エラーが発生しました。"})
			return
		}
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
	err := h.habitRepo.Delete(targetHabitId)

	if err != nil {
		log.Printf("[ERROR] %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "エラーが発生しました。"})
		return
	}

	// 今日のdaily-trackを取得
	todayString := time.Now().Format(`2006-01-02`) // YYYY-MM-DD
	todaysTrack, err := h.dailyTrackRepo.FindDailyTrack(userId, todayString)
	if err != nil && !errors.Is(err, common.ErrNotFound) {
		log.Printf("[ERROR] %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "エラーが発生しました。"})
		return
	}

	if errors.Is(err, common.ErrNotFound) {
		c.JSON(http.StatusOK, gin.H{"message": "success"})
		return
	}

	// 今日のdaily-trackがあれば習慣を削除
	for index, habitStatus := range todaysTrack.HabitStatuses {
		if habitStatus.HabitId == targetHabitId && !habitStatus.IsDone {
			// 削除対象の習慣が完了していなければ削除
			todaysTrack.HabitStatuses = append(todaysTrack.HabitStatuses[:index], todaysTrack.HabitStatuses[index+1:]...)

			// 永続化
			err = h.dailyTrackRepo.UpdateHabitStatuses(todaysTrack)
			if err != nil {
				log.Printf("[ERROR] %v", err)
				c.JSON(http.StatusInternalServerError, gin.H{"message": "エラーが発生しました。"})
				return
			}

			break
		}
	}

	c.JSON(http.StatusOK, gin.H{"message": "success"})
}
