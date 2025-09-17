package utils

import "github.com/gin-gonic/gin"

// Gin Context からユーザーIDを取得する
func GetUserIdFromContext(c *gin.Context) string {
	// NOTE: user_idの検証はミドルウェアに実装
	var userId string
	loginedUserId, exists := c.Get("user_id")
	if !exists {
		return ""
	}

	userId = loginedUserId.(string)
	return userId
}
