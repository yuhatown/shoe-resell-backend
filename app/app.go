package app

import (
	"shoe-resell-backend/model"
	"time"

	"github.com/gin-gonic/gin"
)

func PostUser(c *gin.Context) {

	info := model.User {Name: "test", Email: 20, Password: "asd23", Created_at: time.Time{}, Updated_at: &time.Time{} }
	info.PostUserModel()

	
	// c.JSON(200, gin.H{
	// 	"result": "ok",
	// })


}