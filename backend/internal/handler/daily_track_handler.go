package handler

import (
	"errors"
	"log"
	"net/http"

	"backend/internal/config"
	"backend/internal/domain/common"
	"backend/internal/domain/model/daily_track"
	"backend/internal/domain/model/habit"
	"backend/internal/domain/model/user"
	"backend/internal/domain/repository"
	"backend/internal/utils"

	"github.com/gin-gonic/gin"
)

type DailyTrackHandler struct {
	dailyTrackRepo repository.DailyTrackRepository
	habitRepo      repository.HabitRepository
	userRepo       repository.UserRepository
}

func NewDailyTrackHandler(
	dailyTrackRepository repository.DailyTrackRepository,
	habitRepository repository.HabitRepository,
	userRepository repository.UserRepository,
) *DailyTrackHandler {
	return &DailyTrackHandler{
		dailyTrackRepo: dailyTrackRepository,
		habitRepo:      habitRepository,
		userRepo:       userRepository,
	}
}

func (h *DailyTrackHandler) GetDailyTrack(c *gin.Context) {
	dateParam := c.Param("date")

	// idが空文字列の場合のチェック
	if dateParam == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "日付を指定してください。"})
		return
	}

	userId := utils.GetUserIdFromContext(c)

	todaysTrack, err := h.dailyTrackRepo.FindDailyTrack(userId, dateParam)
	if err != nil && !errors.Is(err, common.ErrNotFound) {
		log.Printf("[ERROR] %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "エラーが発生しました。"})
		return
	}

	// 存在しなかったら作成
	if errors.Is(err, common.ErrNotFound) {
		// 習慣一覧を取得
		var habits []*habit.Habit
		habits, err = h.habitRepo.FetchAll(userId)
		if err != nil {
			log.Printf("[ERROR] %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"message": "エラーが発生しました。"})
			return
		}

		// 習慣一覧を使ってステータスの配列を作成
		var habitStatuses []*daily_track.HabitStatus
		for _, habit := range habits {
			habitStatus := &daily_track.HabitStatus{
				HabitId:   habit.Id,
				HabitName: habit.Name,
				IsDone:    false,
			}
			habitStatuses = append(habitStatuses, habitStatus)
		}

		newDailyTrack := daily_track.DailyTrack{
			UserId:        userId,
			Date:          dateParam,
			HabitStatuses: habitStatuses,
		}

		todaysTrack, err = h.dailyTrackRepo.RegisterDailyTrack(&newDailyTrack)
		if err != nil {
			log.Printf("[ERROR] %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"message": "エラーが発生しました。"})
			return
		}
	}

	c.JSON(http.StatusOK, todaysTrack)
}

// TODO: requestパッケージ作成
type updateDoneDailyTrackRequest struct {
	Date    string `json:"date"   binding:"required"`
	HabitId string `json:"habit_id" binding:"required"`
}

func (h *DailyTrackHandler) UpdateDoneDailyTrack(c *gin.Context) {
	// バリデーション
	var updateDoneDailyTrackRequest updateDoneDailyTrackRequest
	if err := c.ShouldBindJSON(&updateDoneDailyTrackRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request body"})
		return
	}

	// ユーザーID取得
	userId := utils.GetUserIdFromContext(c)

	// todaysTrack 取得
	todaysTrack, err := h.dailyTrackRepo.FindDailyTrack(userId, updateDoneDailyTrackRequest.Date)
	if err != nil {
		log.Printf("[ERROR] %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "エラーが発生しました。"})
		return
	}

	// todaysTrack.IsDone 更新
	for _, habitStatus := range todaysTrack.HabitStatuses {
		if habitStatus.HabitId == updateDoneDailyTrackRequest.HabitId {
			habitStatus.IsDone = true
		}
	}

	// dailyTrack更新して保存
	err = h.dailyTrackRepo.UpdateHabitStatuses(todaysTrack)
	if err != nil {
		log.Printf("[ERROR] %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "エラーが発生しました。"})
		return
	}

	// ユーザー情報取得
	var user *user.User
	user, err = h.userRepo.Find(userId)
	if err != nil {
		log.Printf("[ERROR] %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "エラーが発生しました。"})
		return
	}

	// point 加算
	points := user.Points + config.PointsForHabitDone
	err = h.userRepo.UpdatePoints(userId, points)
	if err != nil {
		log.Printf("[ERROR] %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "エラーが発生しました。"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success"})
}
