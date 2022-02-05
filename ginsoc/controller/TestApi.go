package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Ping(ctx *gin.Context) {
	// http://localhost:8080/api/test/ping?name=11
	name := ctx.DefaultQuery("name", "")
	fmt.Println(name)

	ctx.JSON(200, gin.H{
		"message": "pong",
	})
}

type msg struct {
	Name    string
	Message string
	Age     int
}

func PingJson(ctx *gin.Context) {
	data := msg{
		Name:    "乔四美",
		Message: "七七",
		Age:     77,
	}
	ctx.JSON(200, data)
}
