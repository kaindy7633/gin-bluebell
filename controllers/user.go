package controllers

import (
	"errors"
	"fmt"
	"gin-bluebell/common"
	"gin-bluebell/dao/mysql"
	"gin-bluebell/logic"
	"gin-bluebell/models"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"

	"gin-bluebell/utils"
)

// SingUpHandler 处理注册请求
// @Summary 处理注册请求
// @Description 处理注册请求
// @Tags 用户相关
// @Accept application/json
// @Produce appliction/json
// @Param Authorization header string false "Bearer 用户令牌"
// @Param object query models.ParamSignUp true "注册参数"
// @Security ApiKeyAuth
// @Success 200 {object} models.CommonResponse
// @router /signup [post]
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
			common.ResponseError(c, common.CodeInvalidParam)
			return
		}
		// validator.ValidationErrors 类型错误则进行翻译
		common.ResponseErrorWithMsg(
			c,
			common.CodeInvalidParam,
			utils.RemoveTopStruct(errs.Translate(utils.Trans)),
		)
		return
	}

	// 2. 业务处理
	if err := logic.SignUp(p); err != nil {
		zap.L().Error("logic.Signup failed", zap.Error(err))
		if errors.Is(err, mysql.ErrorUserExist) {
			// 用户已存在
			common.ResponseError(c, common.CodeUserExist)
			return
		}
		common.ResponseError(c, common.CodeServerBusy)
		return
	}

	// 3. 返回响应
	common.ResponseSuccess(c, nil)
}

// 用户登录
func LoginHandler(c *gin.Context) {
	// 1. 获取请求参数及参数校验
	p := new(models.ParamLogin)
	if err := c.ShouldBindJSON(&p); err != nil {
		// 请求参数有误
		zap.L().Error("Login with invalid param", zap.Error(err))
		// 判断错误是否为 validator.ValidationErrors 类型
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			common.ResponseError(c, common.CodeInvalidParam)
			return
		}
		common.ResponseErrorWithMsg(
			c,
			common.CodeInvalidParam,
			utils.RemoveTopStruct(errs.Translate(utils.Trans)),
		)
		return
	}
	// 2. 业务逻辑处理
	user, err := logic.Login(p)
	if err != nil {
		zap.L().Error("logic.Login failed,", zap.Error(err))
		if errors.Is(err, mysql.ErrorUserNotExist) {
			common.ResponseError(c, common.CodeUserNotExist)
			return
		}
		common.ResponseError(c, common.CodeInvalidPassword)
		return
	}
	// 3. 返回响应
	common.ResponseSuccess(c, gin.H{
		"user_id":       fmt.Sprintf("%d", user.UserID), // 这里的值在前端会出现失真的问题
		"username":      user.Username,
		"access_token":  user.AccessToken,
		"refresh_token": user.RefreshToken,
	})
}
