package logic

import (
	"gin-bluebell/dao/mysql"
	"gin-bluebell/models"
	"gin-bluebell/pkg/snowflake"
)

func CreatePost(p *models.Post) (err error) {
	// 1. 使用雪花算法创建 post id
	p.ID = snowflake.GenID()
	// 2. 保存到数据库并返回
	return mysql.CreatePost(p)
}

func GetPostDetailById(pid int64) (data *models.Post, err error) {
	return mysql.GetPostDetailById(pid)
}
