package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()

	app.GET("/v1", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"ok": true,
		})
	})

	app.Run("localhost:3000")
}
