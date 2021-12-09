package middleware

import (
	"gin-bluebell/common"
	"gin-bluebell/pkg/jwt"
	"strings"

	"github.com/gin-gonic/gin"
)

const CtxUserIDKey = "userID"

// JWTAuthMiddleware 基于JWT的认证中间件
func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		// 客户端携带Token有三种方式 1.放在请求头 2.放在请求体 3.放在URI
		// 这里假设Token放在Header的Authorization中，并使用Bearer开头
		// Authorization: Bearer: token值
		// 这里的具体实现方式要根据你的实际业务决定
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			common.ResponseError(c, common.CodeMissingToken)
			c.Abort()
			return
		}

		// 按空格进行分割拿到token值
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			common.ResponseError(c, common.CodeInvalidToken)
			c.Abort()
			return
		}

		// parts[1] 是获取到的 tokenString，我们使用之前定义好的解析JWT的函数来解析它
		mc, err := jwt.ParseToken(parts[1])
		if err != nil {
			common.ResponseError(c, common.CodeInvalidToken)
			c.Abort()
			return
		}

		// 将当前请求的 userID 信息保存到请求上下文中
		// 后续的处理请求函数中，可以用 c.Get(CtxUserIDKey) 来获取当前请求的用户信息
		c.Set(CtxUserIDKey, mc.UserID)
		c.Next()
	}
}
