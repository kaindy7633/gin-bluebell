package routes

import (
	"gin-bluebell/controllers"
	"gin-bluebell/logger"
	"gin-bluebell/middleware"
	"net/http"

	_ "gin-bluebell/docs"

	gs "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	// swagger route
	r.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))
	/**
	 * gin-swagger同时还提供了DisablingWrapHandler函数，
	 * 方便我们通过设置某些环境变量来禁用Swagger。例如：
	 * r.GET("/swagger/*any", gs.DisablingWrapHandler(swaggerFiles.Handler, "NAME_OF_ENV_VARIABLE"))
	 * 此时如果将环境变量NAME_OF_ENV_VARIABLE设置为任意值，则/swagger/*any将返回404响应
	 */

	// 创建 v1 版本的路由
	v1 := r.Group("/api/v1")

	// 注册业务路由
	v1.POST("/signup", controllers.SignUpHandler)
	// 登录路由
	v1.POST("/login", controllers.LoginHandler)
	// 应用 JWT 认证中间件
	v1.Use(middleware.JWTAuthMiddleware())
	{
		// 社区分类
		v1.GET("/community", controllers.CommunityHandler)
		v1.GET("/community/:id", controllers.CommunityDetailHandler)

		// 帖子
		v1.POST("/post", controllers.CreatePostHandler)
		v1.GET("/post/:id", controllers.GetPostDetailHandler)
		v1.GET("/posts", controllers.GetPostListHandler)
		// 对获取帖子的接口进行扩展，加入根据参数进行筛选排序的功能
		// 如：按评分排序获取、按创建时间获取
		v1.GET("/posts_sort", controllers.GetPostListHandlerByParams)

		// 投票
		v1.POST("/vote", controllers.PostVoteController)
	}

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})
	})

	return r
}
