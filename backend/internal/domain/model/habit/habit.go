package habit

type Habit struct {
	Id     string `json:"id"`
	UserId string `json:"user_id"`
	Name   string `json:"name"`
}
