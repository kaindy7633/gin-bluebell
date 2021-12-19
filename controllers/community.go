package controllers

import (
	"gin-bluebell/common"
	"gin-bluebell/logic"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// ---- 社区相关 ----

func CommunityHandler(c *gin.Context) {
	// 查询到所有的社区（community_id, community_name）以列表的形式返回
	data, err := logic.GetCommunityList()
	if err != nil {
		zap.L().Error("logic.GetCommunityList() faild", zap.Error(err))
		common.ResponseError(c, common.CodeServerBusy)
		return
	}
	common.ResponseSuccess(c, data)
}

// CommunityDetailHandler 社区分类详情
func CommunityDetailHandler(c *gin.Context) {
	// 1.获取社区id
	idStr := c.Param("id")                     // string 类型的id
	id, err := strconv.ParseInt(idStr, 10, 64) // 转换为 int
	if err != nil {
		common.ResponseError(c, common.CodeInvalidParam)
		return
	}

	// 根据id查询详情
	data, err := logic.GetCommunityDetail(id)
	if err != nil {
		zap.L().Error("logic.GetCommunityDetail failed", zap.Error(err))
		common.ResponseError(c, common.CodeServerBusy)
		return
	}
	common.ResponseSuccess(c, data)
}
