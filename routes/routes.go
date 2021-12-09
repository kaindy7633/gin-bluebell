package routes

import (
	"gin-bluebell/controllers"
	"gin-bluebell/logger"
	"gin-bluebell/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	// 注册业务路由
	r.POST("/signup", controllers.SignUpHandler)
	// 登录路由
	r.POST("/login", controllers.LoginHandler)

	// 测试jwt
	r.GET("/ping", middleware.JWTAuthMiddleware(), func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	return r
}
