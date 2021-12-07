package mysql

import (
	"database/sql"
	"errors"
	"gin-bluebell/models"
	"gin-bluebell/utils"
)

// CheckUserExist 判断用户是否存在
func CheckUserExist(username string) (err error) {
	sqlStr := `select count(user_id) from user where username = ?`
	var count int
	if err := db.Get(&count, sqlStr, username); err != nil {
		return err
	}
	if count > 0 {
		return errors.New("用户已存在")
	}
	return
}

// InsertUser 向数据库中插入一条心的用户记录
func InsertUser(user *models.User) (err error) {
	// 对密码进行加密
	user.Password = utils.EncryptString(user.Password)
	// 执行 SQL 语句入库
	sqlStr := `insert into user (user_id, username, password) values(?, ?, ?)`
	_, err = db.Exec(sqlStr, user.UserID, user.Username, user.Password)
	return
}

// Login
func Login(user *models.User) (err error) {
	_password := user.Password
	sqlStr := `select user_id, username, password from user where username = ?`
	err = db.Get(user, sqlStr, user.Username)
	// 判断用户是否存在
	if err == sql.ErrNoRows {
		return errors.New("用户不存在")
	}
	if err != nil {
		return err
	}
	// 判断密码是否正确
	if user.Password != utils.EncryptString(_password) {
		return errors.New("密码错误")
	}
	return

}
