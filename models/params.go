package models

// 定义请求的参数结构体

type ParamSignUp struct {
	Username   string `json:"username" binding:"required"`
	Password   string `json:"password" binding:"required"`
	RePassword string `json:"re_password" binding:"required,eqfield=Password"`
}

type ParamLogin struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// 投票数据
type ParamVoteData struct {
	// UserID 从请求中获取当前用户
	PostID    string `json:"post_id" binding:"required"`       // 帖子 ID
	Direction int8   `json:"direction" binding:"oneof=1 0 -1"` // 赞成票（1）和是反对票（-1） 取消投票（0） 方向
}
