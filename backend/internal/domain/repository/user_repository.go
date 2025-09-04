package repository

// UserRepository はユーザー関連のDB操作を抽象化します
type UserRepository interface {
	FindAll() ([]User, error)
	// ...
}
