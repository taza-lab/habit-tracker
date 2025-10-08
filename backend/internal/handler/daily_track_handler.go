package handler

// handler規約
// フロントで表示するメッセージはここに定義
// サービスからのエラーはlog.Printf("[ERROR] ~")でそのまま出力

import (
	"log"
	"net/http"

	"backend/internal/domain/service"
	"backend/internal/utils"

	"github.com/gin-gonic/gin"
)

type DailyTrackHandler struct {
	dailyTrackService service.DailyTrackService
}

func NewDailyTrackHandler(dailyTrackService service.DailyTrackService) *DailyTrackHandler {
	return &DailyTrackHandler{
		dailyTrackService: dailyTrackService,
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
	todaysTrack, err := h.dailyTrackService.GetDailyTrack(c.Request.Context(), userId, dateParam)
	if err != nil {
		log.Printf("[ERROR] %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "エラーが発生しました。"})
		return
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
	err := h.dailyTrackService.UpdateDoneDailyTrack(c.Request.Context(), userId, updateDoneDailyTrackRequest.Date, updateDoneDailyTrackRequest.HabitId)

	if err != nil {
		log.Printf("[ERROR] %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "エラーが発生しました。"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success"})
}
