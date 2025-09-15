package daily_track

type DailyTrack struct {
	UserId        string        `json:"user_id"`
	Date          string        `json:"date"`
	HabitStatuses []HabitStatus `json:"habit_statuses"`
}
