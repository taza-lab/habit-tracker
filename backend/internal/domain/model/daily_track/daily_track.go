package daily_track

type DailyTrack struct {
	Id            string         `json:"id"`
	UserId        string         `json:"user_id"`
	Date          string         `json:"date"`
	HabitStatuses []*HabitStatus `json:"habit_statuses"`
}
