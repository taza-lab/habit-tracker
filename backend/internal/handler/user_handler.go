package handler

import (
	"net/http"
	"time"
	"os"
	"log"

    "github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	userModel "backend/internal/domain/model/user"
	"backend/internal/domain/repository"
	"backend/internal/domain/common"
)

type UserHandler struct {
	userRepo repository.UserRepository
}

func NewUserHandler(repo repository.UserRepository) *UserHandler {
	return &UserHandler{
		userRepo: repo,
	}
}

type SignUpRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	ConfirmPassword string `json:"confirm_password" binding:"required"`
}

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h *UserHandler) SignUp(c *gin.Context) {
	var signUpRequest SignUpRequest

	// リクエスト内容の検証・構造体バインド
	if err := c.ShouldBindJSON(&signUpRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	if signUpRequest.Password != signUpRequest.ConfirmPassword {
        c.JSON(http.StatusBadRequest, gin.H{"error": "確認用パスワードが一致しません。"})
        return
    }

	// 同一usernameが登録済みかどうかのチェック
	_, err := h.userRepo.FindByUserName(signUpRequest.Username)

	if err != nil && err != common.ErrNotFound {
		log.Printf("failed to find user: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "エラーが発生しました。"})
		return
	}
	if err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "使用できないユーザーネームです。"})
		return
	}
	
	// 登録
	user := userModel.User{Username: signUpRequest.Username, Password:signUpRequest.Password, Points: 0}
	result, err := h.userRepo.Register(&user)
	result.Password = ""

	c.JSON(http.StatusOK, gin.H{
		"user": result,
	})
}

func (h *UserHandler) Login(c *gin.Context) {
	var jwtSecret = []byte(os.Getenv("JWT_SECRET_KEY"))
	var loginRequest LoginRequest

	// リクエスト内容の検証・構造体バインド
	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	user, err := h.userRepo.FindByUserName(loginRequest.Username)

	if err != nil {
		if err == common.ErrNotFound {
			c.JSON(http.StatusBadRequest, gin.H{"error": "ユーザーが登録されていません。"})
			return
		}

		log.Printf("failed to find user: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "エラーが発生しました。"})
			return
	}

	// パスワードチェック
	// bcrypt.CompareHashAndPasswordが保存されているハッシュ値とユーザーが入力したパスワードが一致するかを検証
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginRequest.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "パスワードが正しくありません。"})
		return
	}
	user.Password = ""

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
		log.Printf("Could not create token: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "エラーが発生しました。"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": tokenString,
		"user": user,
	})
}

func (h *UserHandler) GetUser(c *gin.Context) {
	var data = userModel.User{Points: 20}

	c.JSON(http.StatusOK, data)
}

