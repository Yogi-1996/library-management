package main

import (
	"fmt"
	"net/http"

	"github.com/Yogi-1996/library-management/internal/config"
	"github.com/gin-gonic/gin"
)

func main() {

	config := config.LoadConfig()
	fmt.Print(config)

	router := gin.Default()

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Library managment is ready",
			"status":  "sucess",
		})
	})
	router.Run(":" + config.App.Port)
}
