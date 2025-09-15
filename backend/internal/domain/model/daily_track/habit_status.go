package daily_track

type HabitStatus struct {
    HabitId   string `json:"habit_id"`
    HabitName string `json:"habit_name"`
    IsDone    bool   `json:"is_done"`
}
