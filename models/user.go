package models

type User struct {
	UserID       int64  `db:"user_id"`
	Username     string `db:"username"`
	Password     string `db:"password"`
	AccessToken  string
	RefreshToken string
}
