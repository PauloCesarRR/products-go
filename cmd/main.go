package main

import "github.com/gin-gonic/gin"

func main() {
	server := gin.Default()

	server.GET("ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	err := server.Run(":8080")
	if err != nil {
		return
	}
}
