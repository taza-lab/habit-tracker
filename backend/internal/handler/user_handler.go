package handler

import (
	"net/http"

    "github.com/gin-gonic/gin"
	"backend/internal/domain/model"
)

func GetUser(c *gin.Context) {
	var data = model.User{Point: 20}

	c.JSON(http.StatusOK, data)
}