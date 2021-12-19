package mysql

import (
	"database/sql"
	"gin-bluebell/models"

	"go.uber.org/zap"
)

// GetCommunityList 查询社区分类列表
func GetCommunityList() (communityList []*models.Community, err error) {
	sqlStr := "SELECT community_id, community_name from community"
	if err := db.Select(&communityList, sqlStr); err != nil {
		if err == sql.ErrNoRows {
			zap.L().Warn("no rows")
			// err = nil
		}
	}
	return
}

// GetCommunityDetailByID 根据 ID 查询社区分类详情
func GetCommunityDetailByID(id int64) (communityDetail *models.CommunityDetail, err error) {
	communityDetail = new(models.CommunityDetail)
	sqlStr := "SELECT community_id, community_name, introduction, create_time from community where community_id = ?"
	if err = db.Get(communityDetail, sqlStr, id); err != nil {
		if err == sql.ErrNoRows {
			err = ErrorInvalidID
		}
	}
	return communityDetail, err
}
