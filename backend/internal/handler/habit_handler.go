package handler

import (
	"net/http"
	"time"

    "github.com/gin-gonic/gin"
)

// メモ
// gin.H = map[string]interface{}

// TODO: モデル作成
type Habit struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type HabitStatus struct {
    Habit  Habit `json:"habit"`
    IsDone bool  `json:"isDone"`
}

type DailyTrack struct {
	Date   string        `json:"date"`
	Habits []HabitStatus `json:"habits"`
}

func GetHabitList(c *gin.Context) {
	var data = []Habit{
		{Id: 1, Name: "朝ヨガ"},
		{Id: 2, Name: "勉強"},
	}

	c.JSON(http.StatusOK, data)
}

func GetHabit(c *gin.Context) {
	var data = Habit{Id: 1, Name: "朝ヨガ"}

	c.JSON(http.StatusOK, data)
}

func RegisterHabit(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "success", "id": time.Now().Format("20060102150405")})
}

func UpdateHabit(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "success"})
}

func DeleteHabit(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "success"})
}

func GetDailyTrack(c *gin.Context) {
	var data = DailyTrack{
		Date: c.Param("date"),
		Habits: []HabitStatus{
			{Habit: Habit{Id: 1, Name: "朝ヨガ"}, IsDone: false},
			{Habit: Habit{Id: 2, Name: "勉強"}, IsDone: false},
		},
	}

	c.JSON(http.StatusOK, data)
}

func UpdateDoneDailyTrack(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "success"})
}