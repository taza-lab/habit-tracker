package handler

import (
	"net/http"
	"time"
	"os"

    "github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	userModel "backend/internal/domain/user"
)

// 認証情報 TODO: リクエスト定義する
type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Login(c *gin.Context) {
	var jwtSecret = []byte(os.Getenv("JWT_SECRET_KEY"))
	var loginRequest LoginRequest

	// リクエスト内容の検証・構造体バインド
	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// TODO: DBからユーザー取得
	userId := "123ABC"

	// TODO: DBからポイント取得
	points := 20

	user := userModel.User{Id: userId, Username: "testuser", Points: points}

	// 認証 TODO: DBから取得、パスワード復号化
	if loginRequest.Username != user.Username || loginRequest.Password != "password" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	// JWTトークンの生成
	expirationTime := time.Now().Add(24 * time.Hour) // 有効期限は1日
	claims := &userModel.Claims{
		UserId: user.Id,
		Username: loginRequest.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": tokenString,
		"user": user,
	})
}

func GetUser(c *gin.Context) {
	var data = userModel.User{Points: 20}

	c.JSON(http.StatusOK, data)
}