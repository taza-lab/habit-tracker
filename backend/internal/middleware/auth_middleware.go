package middleware

import (
	"net/http"
	"os"

    "github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"backend/internal/domain/model/user"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
			c.Abort()
			return
		}

		// トークンから "Bearer " を除去
		if len(tokenString) > 7 && tokenString[:7] == "Bearer " {
			tokenString = tokenString[7:]
		}

		var jwtSecret = []byte(os.Getenv("JWT_SECRET_KEY"))
		claims := &user.Claims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		})

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		if !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token is not valid"})
			c.Abort()
			return
		}

		// 認証が成功したら、ユーザー名をコンテキストに保存
		c.Set("username", claims.Username)
		c.Next()
	}
}
