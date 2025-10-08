package handler

// handler規約
// フロントで表示するメッセージはここに定義
// サービスからのエラーはlog.Printf("[ERROR] ~")でそのまま出力

import (
	"errors"
	"log"
	"net/http"

	"backend/internal/domain/common"
	"backend/internal/domain/service"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

type SignUpRequest struct {
	Username        string `json:"username" binding:"required"`
	Password        string `json:"password" binding:"required"`
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
		c.JSON(http.StatusBadRequest, gin.H{"error": "リクエストが不正です。"})
		return
	}

	if signUpRequest.Password != signUpRequest.ConfirmPassword {
		c.JSON(http.StatusBadRequest, gin.H{"error": "確認用パスワードが一致しません。"})
		return
	}

	// サインアップサービス実行
	result, err := h.userService.SignUp(c.Request.Context(), signUpRequest.Username, signUpRequest.Password)

	if errors.Is(err, common.ErrAlreadyExists) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "使用できないユーザーネームです。"})
		return
	}
	if err != nil {
		log.Printf("[ERROR] %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "エラーが発生しました。"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user": result,
	})
}

func (h *UserHandler) Login(c *gin.Context) {
	var loginRequest LoginRequest

	// リクエスト内容の検証・構造体バインド
	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "リクエストが不正です。"})
		return
	}

	// ログインサービス実行
	user, tokenString, err := h.userService.Login(c.Request.Context(), loginRequest.Username, loginRequest.Password)

	if err != nil {
		if err == common.ErrNotFound {
			c.JSON(http.StatusBadRequest, gin.H{"error": "ユーザーが登録されていません。"})
			return
		}

		if err == common.ErrPasswordMismatch {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "パスワードが正しくありません。"})
			return
		}

		log.Printf("[ERROR] %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "エラーが発生しました。"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": tokenString,
		"user":  user,
	})
}
