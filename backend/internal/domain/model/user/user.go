package user

type User struct {
	Id string `json:"id"`
	Username string `json:"username"`
	Points int `json:"points"`
}