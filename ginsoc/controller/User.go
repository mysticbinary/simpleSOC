package controller

import (
	"fmt"
	"ginsoc/models"

	"github.com/gin-gonic/gin"
)

func FindAll(ctx *gin.Context) {
	data := models.User{
		Id:      1,
		Name:    "张三",
		Section: "security",
		Role:    "admin",
	}
	ctx.JSON(200, data)
}

func ChangePwd(ctx *gin.Context) {
	message := ctx.PostForm("message")
	fmt.Println("changePwd message:", message)

	name := ctx.DefaultQuery("username", "")
	fmt.Println(name)

	ctx.JSON(200, gin.H{
		"password": "123456",
	})
}
