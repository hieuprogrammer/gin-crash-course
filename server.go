package main

import "github.com/gin-gonic/gin"

func main() {
	server := gin.Default()

	server.GET("/test", func(context *gin.Context) {
		context.JSON(200, gin.H {
			"message": "Gin is awesome! :D",
		})
	})

	server.Run(":8080")
}