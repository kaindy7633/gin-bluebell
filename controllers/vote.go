package controllers

import (
	"errors"
	"gin-bluebell/common"
	"gin-bluebell/logic"
	"gin-bluebell/middleware"
	"gin-bluebell/models"
	"gin-bluebell/utils"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

// 投票
type VoteData struct {
	PostID    int64 `json:"post_id,string"`   // 帖子 ID
	Direction int   `json:"direction,string"` // 赞成票（1）和是反对票（-1） 方向
}

func PostVoteController(c *gin.Context) {
	// 参数校验
	p := new(models.ParamVoteData)
	if err := c.ShouldBindJSON(p); err != nil {
		errs, ok := err.(validator.ValidationErrors) // 类型断言

		if !ok {
			common.ResponseError(c, common.CodeInvalidParam)
			return
		}
		errData := utils.RemoveTopStruct(errs.Translate(utils.Trans)) // 翻译并去除错误提示中的结构体标识
		common.ResponseErrorWithMsg(c, common.CodeInvalidParam, errData)
		return
	}

	// 获取用户id
	userID, ok := c.Get(middleware.CtxUserIDKey)
	if !ok {
		zap.L().Error("c.Get(middleware.CtxUserIDKey) failed", zap.Error(errors.New("获取用户ID失败")))
		common.ResponseError(c, common.CodeServerBusy)
		return
	}

	if err := logic.VoteForPost(userID.(int64), p); err != nil {
		zap.L().Error("logic.VoteForPost failed", zap.Error(err))
		common.ResponseError(c, common.CodeServerBusy)
		return
	}
	common.ResponseSuccess(c, nil)

}
