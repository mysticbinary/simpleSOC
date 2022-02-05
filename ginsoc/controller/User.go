package controller

import (
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
