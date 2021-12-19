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

	// 创建 v1 版本的路由
	v1 := r.Group("/api/v1")

	// 注册业务路由
	v1.POST("/signup", controllers.SignUpHandler)
	// 登录路由
	v1.POST("/login", controllers.LoginHandler)
	// 应用 JWT 认证中间件
	v1.Use(middleware.JWTAuthMiddleware())
	{
		v1.GET("/community", controllers.CommunityHandler)
		v1.GET("/community/:id", controllers.CommunityDetailHandler)
	}

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})
	})

	return r
}
