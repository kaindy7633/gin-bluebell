package mysql

import (
	"database/sql"
	"gin-bluebell/models"
	"gin-bluebell/utils"
)

// var (
// 	ErrorUserExist       = errors.New("用户已存在")
// 	ErrorUserNotExist    = errors.New("用户不存在")
// 	ErrorInvalidPassword = errors.New("用户名或密码错误")
// )

// CheckUserExist 判断用户是否存在
func CheckUserExist(username string) (err error) {
	sqlStr := `select count(user_id) from user where username = ?`
	var count int
	if err := db.Get(&count, sqlStr, username); err != nil {
		return err
	}
	if count > 0 {
		return ErrorUserExist
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
		return ErrorUserNotExist
	}
	if err != nil {
		return err
	}
	// 判断密码是否正确
	if user.Password != utils.EncryptString(_password) {
		return ErrorInvalidPassword
	}
	return

}

// 根据id获取用户信息
func GetUserById(uid int64) (user *models.User, err error) {
	user = new(models.User)
	sqlStr := `SELECT user_id, username FROM user WHERE user_id = ?`
	err = db.Get(user, sqlStr, uid)
	return
}
