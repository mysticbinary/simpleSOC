package controller

import (
	"fmt"
	"ginsoc/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Tab struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func FindAll(ctx *gin.Context) {
	inTab := []models.User{}
	for i := 0; i < 6; i++ {
		tab := models.User{}
		tab.Id = i
		tab.Name = "name" + strconv.Itoa(i)
		tab.Role = "admin"
		tab.Section = "security"

		inTab = append(inTab, tab)
	}
	data := gin.H{
		"code": 20000,
		"data": inTab,
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
