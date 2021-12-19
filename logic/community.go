package logic

import (
	"gin-bluebell/dao/mysql"
	"gin-bluebell/models"
)

func GetCommunityList() ([]*models.Community, error) {
	// 查询数据库，获取所有的 community，并返回
	return mysql.GetCommunityList()
}

func GetCommunityDetail(id int64) (*models.CommunityDetail, error) {
	return mysql.GetCommunityDetailByID(id)
}
