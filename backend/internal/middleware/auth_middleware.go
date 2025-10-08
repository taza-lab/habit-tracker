package middleware

import (
	"net/http"
	"os"
	"strings"

	"backend/internal/domain/model/user"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

var jwtSecret []byte

// 初期化関数
func InitJWTSecret() {
	// NOTE: main()で.envファイルを読み込んだ後に実行
	jwtSecret = []byte(os.Getenv("JWT_SECRET_KEY"))
	if len(jwtSecret) == 0 {
		panic("JWT_SECRET_KEY environment variable is not set")
	}
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
			c.Abort()
			return
		}

		// プレフィックスのチェックとトークンの抽出
		const bearerPrefix = "Bearer "
		tokenString := strings.TrimPrefix(authHeader, bearerPrefix)
		if tokenString == authHeader { // TrimPrefixが何も変更しなかった場合
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token format"})
			c.Abort()
			return
		}

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

		// 認証成功、ユーザーIDをコンテキストに保存
		c.Set("user_id", claims.UserId)
		c.Next()
	}
}
