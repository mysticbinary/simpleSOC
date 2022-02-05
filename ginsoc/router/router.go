package router

import (
	"ginsoc/controller"

	"github.com/gin-gonic/gin"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	// test
	r.GET("/api/test/ping", controller.Ping)
	r.GET("/api/test/pingJson", controller.PingJson)

	// login
	r.POST("/api/user/login", controller.Login)
	r.POST("/api/user/logout", controller.Logout)
	r.GET("/api/user/info", controller.Info)

	// user
	r.GET("/api/user/findAll", controller.FindAll)

	return r
}
