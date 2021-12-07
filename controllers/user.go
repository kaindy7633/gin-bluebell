package controllers

import (
	"gin-bluebell/logic"
	"gin-bluebell/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"

	"gin-bluebell/utils"
)

// SingUpHandler 处理注册请求的函数
func SignUpHandler(c *gin.Context) {
	// 1. 获取参数和参数校验
	p := new(models.ParamSignUp)
	if err := c.ShouldBindJSON(&p); err != nil {
		// 请求参数有误，直接返回响应
		zap.L().Error("Signup with invalid param", zap.Error(err))
		// 获取 validator.ValidationErrors 类型的 errors
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			// 非 validator.ValidationErrors 类型错误直接返回
			c.JSON(http.StatusOK, gin.H{
				"msg": err.Error(),
			})
			return
		}
		// validator.ValidationErrors 类型错误则进行翻译
		c.JSON(http.StatusOK, gin.H{
			"msg": utils.RemoveTopStruct(errs.Translate(utils.Trans)),
		})
		return
	}

	// 2. 业务处理
	if err := logic.SignUp(p); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": "注册失败",
		})
	}

	// 3. 返回响应
	c.JSON(http.StatusOK, gin.H{
		"msg": "success",
	})
}
