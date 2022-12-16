package main

import (
	"github.com/gin-gonic/gin"
	"shoe-resell-backend/app"
)

func main() {
	router := gin.Default()

	v1 := router.Group("/user")
	{
		v1.POST("/login", app.PostUser)
		// v1.PUT("/submit", )
		// v1.GET("/read", )
		// v1.DELETE("/submit", )
	}

	router.Run(":8080")
}