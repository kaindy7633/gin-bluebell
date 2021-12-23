package logic

import (
	"gin-bluebell/dao/mysql"
	"gin-bluebell/models"
	"gin-bluebell/pkg/snowflake"

	"go.uber.org/zap"
)

func CreatePost(p *models.Post) (err error) {
	// 1. 使用雪花算法创建 post id
	p.ID = snowflake.GenID()
	// 2. 保存到数据库并返回
	return mysql.CreatePost(p)
}

func GetPostDetailById(pid int64) (data *models.ApiPostDetail, err error) {
	// 查询并组合数组
	post, err := mysql.GetPostDetailById(pid)
	if err != nil {
		zap.L().Error("mysql.GetPostDetailById failed", zap.Int64("pid", pid), zap.Error(err))
		return
	}

	// 查询作者信息
	user, err := mysql.GetUserById(post.AuthorID)
	if err != nil {
		zap.L().Error("mysql.GetUserById failed", zap.Int64("author_id", post.AuthorID), zap.Error(err))
		return
	}

	// 查询社区信息
	community, err := mysql.GetCommunityDetailByID(post.CommunityID)
	if err != nil {
		zap.L().Error("mysql.GetCommunityDetailByID failed", zap.Int64("community_id", post.CommunityID), zap.Error(err))
		return
	}

	data = &models.ApiPostDetail{
		AuthorName:      user.Username,
		Post:            post,
		CommunityDetail: community,
	}

	return
}
