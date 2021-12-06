package logic

import (
	"gin-bluebell/dao/mysql"
	"gin-bluebell/models"
	"gin-bluebell/pkg/snowflake"
)

func SignUp(p *models.ParamSignUp) {
	// 1. 判断用户是否存在
	mysql.QueryUserByUsername()
	// 2. 生成 UID
	snowflake.GenID()
	// 3. 保存到数据库
	mysql.InsertUser()
}
