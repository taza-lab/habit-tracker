package user

import "github.com/golang-jwt/jwt/v4"

type Claims struct {
	UserId   string `json:"user_id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// Valid() メソッドを実装して jwt.Claims インターフェースを満たす
// RegisteredClaims の Valid() を呼び出すだけで十分
func (c *Claims) Valid() error {
	return c.RegisteredClaims.Valid()
}