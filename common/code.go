package common

type ResCode int64

const (
	CodeSuccess         ResCode = 200
	CodeInvalidParam    ResCode = 400
	CodeUserExist       ResCode = 10001
	CodeUserNotExist    ResCode = 10002
	CodeInvalidPassword ResCode = 10003
	CodeServerBusy      ResCode = 500

	CodeMissingToken ResCode = 401
	CodeInvalidToken ResCode = 405
)

var codeMsgMap = map[ResCode]string{
	CodeSuccess:         "success",
	CodeInvalidParam:    "请求参数错误",
	CodeUserExist:       "用户已存在",
	CodeUserNotExist:    "用户不存在",
	CodeInvalidPassword: "密码错误",
	CodeServerBusy:      "服务忙",
	CodeMissingToken:    "缺少必要的Token",
	CodeInvalidToken:    "无效的Token",
}

func (rc ResCode) Msg() string {
	msg, ok := codeMsgMap[rc]
	if !ok {
		msg = codeMsgMap[CodeServerBusy]
	}
	return msg
}
