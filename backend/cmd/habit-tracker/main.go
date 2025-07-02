package main

import (
	"net/http"

    "github.com/gin-gonic/gin"
)

func ok(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
            "message": "OKです",
    })
}

func main() {
    router := gin.Default()
    router.GET("/", ok)
    router.Run(":8080")
}
