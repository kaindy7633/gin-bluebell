package models

// 定义请求的参数结构体

const (
	OrderTime  = "time"
	OrderScore = "score"
)

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

// ParamPostList 获取帖子列表 query string 参数
type ParamPostList struct {
	Page     int64  `json:"page" form:"page"`
	PageSize int64  `json:"page_size" form:"page_size"`
	Order    string `json:"order" form:"order" `
}
