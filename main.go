package main

import (
	"net/http"
	
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/recv", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"result": "ok",
		})
	})
	r.Run()
}