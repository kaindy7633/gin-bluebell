package controllers

import (
	"errors"
	"gin-bluebell/common"
	"gin-bluebell/logic"
	"gin-bluebell/middleware"
	"gin-bluebell/models"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// CreatePostHandler 创建新的帖子
func CreatePostHandler(c *gin.Context) {
	// 1. 获取参数及参数校验
	p := new(models.Post)
	if err := c.ShouldBindJSON(p); err != nil {
		common.ResponseError(c, common.CodeInvalidParam)
		return
	}

	// 获取用户ID
	userID, ok := c.Get(middleware.CtxUserIDKey)
	if !ok {
		zap.L().Error("c.Get('CtxUserIDKey') failed", zap.Error(errors.New("获取用户ID失败")))
		common.ResponseError(c, common.CodeServerBusy)
		return
	}
	p.AuthorID = userID.(int64)

	// 2. 创建帖子
	if err := logic.CreatePost(p); err != nil {
		zap.L().Error("logic.CreatePost failed", zap.Error(err))
		common.ResponseError(c, common.CodeServerBusy)
		return
	}
	// 3. 返回响应
	common.ResponseSuccess(c, nil)
}

// GetPostDetailHandler 获取帖子详情
func GetPostDetailHandler(c *gin.Context) {
	// 1. 获取参数 (帖子的id)
	pidStr := c.Param("id")
	pid, err := strconv.ParseInt(pidStr, 10, 64)
	if err != nil {
		zap.L().Error("get post detail with invalid param", zap.Error(err))
		common.ResponseError(c, common.CodeInvalidParam)
		return
	}
	// 2. 根据id获取帖子数据
	data, err := logic.GetPostDetailById(pid)
	if err != nil {
		zap.L().Error("logic.GetPostDetailById failed", zap.Error(err))
		common.ResponseError(c, common.CodeServerBusy)
		return
	}
	// 3. 返回响应
	common.ResponseSuccess(c, data)
}
