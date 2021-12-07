package logic

import (
	"gin-bluebell/dao/mysql"
	"gin-bluebell/models"
	"gin-bluebell/pkg/snowflake"
)

func SignUp(p *models.ParamSignUp) (err error) {
	// 1. 判断用户是否存在
	if err := mysql.CheckUserExist(p.Username); err != nil {
		return err
	}

	// 2. 生成 UID
	userID := snowflake.GenID()
	// 构造一个 User 实例
	user := &models.User{
		UserID:   userID,
		Username: p.Username,
		Password: p.Password,
	}
	// 3. 保存到数据库
	return mysql.InsertUser(user)
}