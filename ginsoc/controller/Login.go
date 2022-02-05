package controller

import "github.com/gin-gonic/gin"

//{"code":20000,"data":{"token":"admin-token"}}
type indata struct {
	Token string `json:"token"`
}

func Login(ctx *gin.Context) {
	data := gin.H{
		"code": 20000,
		"data": indata{
			Token: "admin-token",
		},
	}
	ctx.JSON(200, data)
}

//{"code":20000,"data":"success"}
func Logout(ctx *gin.Context) {
	data := gin.H{
		"code": 20000,
		"data": "success",
	}
	ctx.JSON(200, data)
}

/**
{
  "code": 20000,
  "data": {
    "roles": [
      "admin"
    ],
    "introduction": "I am a super administrator",
    "avatar": "https://wpimg.wallstcn.com/f778738c-e4f8-4870-b634-56703b4acafe.gif",
    "name": "Super Admin"
  }
}
*/
type InfoJson struct {
	Code int `json:"code"`
	Data struct {
		Roles        []string `json:"roles"`
		Introduction string   `json:"introduction"`
		Avatar       string   `json:"avatar"`
		Name         string   `json:"name"`
	} `json:"data"`
}

func Info(ctx *gin.Context) {
	data := InfoJson{
		Code: 20000,
	}
	data.Data.Roles = []string{"admin"}
	data.Data.Introduction = "I am admin."
	data.Data.Name = "amdin"
	data.Data.Avatar = "https://wpimg.wallstcn.com/f778738c-e4f8-4870-b634-56703b4acafe.gif"
	ctx.JSON(200, data)
}
