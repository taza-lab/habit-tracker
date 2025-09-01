package model

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