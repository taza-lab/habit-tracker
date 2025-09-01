package habit_track

type HabitStatus struct {
    Habit  Habit `json:"habit"`
    IsDone bool  `json:"isDone"`
}