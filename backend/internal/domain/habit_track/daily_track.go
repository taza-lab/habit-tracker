package habit_track

type DailyTrack struct {
	Date   string        `json:"date"`
	Habits []HabitStatus `json:"habits"`
}