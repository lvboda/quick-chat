package status

import (
	"github.com/gin-gonic/gin"
)

const (
	SUCCESS = 200
	ERROR   = 500

	ERROR_REQUEST_PARAM = 5001
	ERROR_FILE_PARSE    = 5002
	ERROR_FILE_UPLOAD   = 5003

	// code= 1000... 用户模块的错误
	ERROR_USERNAME_USED    = 1001
	ERROR_PASSWORD_WRONG   = 1002
	ERROR_USER_NOT_EXIST   = 1003
	ERROR_TOKEN_CREATE     = 1004
	ERROR_TOKEN_RUNTIME    = 1005
	ERROR_TOKEN_WRONG      = 1006
	ERROR_TOKEN_TYPE_WRONG = 1007
	ERROR_USER_NO_RIGHT    = 1008
	ERROR_USER_UPDATE      = 1009
	ERROR_USER_DELETE      = 1010
)

var statusMsgMap = map[int]string{
	SUCCESS: "OK",
	ERROR:   "ERROR",

	ERROR_REQUEST_PARAM: "请求参数错误",
	ERROR_FILE_PARSE:    "文件解析失败",
	ERROR_FILE_UPLOAD:   "文件上传失败",

	ERROR_USERNAME_USED:    "用户已存在",
	ERROR_PASSWORD_WRONG:   "用户名或密码错误",
	ERROR_USER_NOT_EXIST:   "用户不存在",
	ERROR_TOKEN_CREATE:     "TOKEN创建失败",
	ERROR_TOKEN_RUNTIME:    "TOKEN已过期 请重新登陆",
	ERROR_TOKEN_WRONG:      "TOKEN不正确 请重新登陆",
	ERROR_TOKEN_TYPE_WRONG: "TOKEN格式错误 请重新登陆",
	ERROR_USER_NO_RIGHT:    "该用户无权限",
	ERROR_USER_UPDATE:      "用户信息修改失败",
	ERROR_USER_DELETE:      "用户注销失败",
}

func GetStatusMsg(status int) string {
	return statusMsgMap[status]
}

func GetResponse(status int, message any, data any) gin.H {
	if v, ok := message.(error); ok {
		message = GetStatusMsg(status) + "--" + v.Error()
	} else if v, ok := message.(string); ok {
		message = GetStatusMsg(status) + "--" + v
	} else if message == nil {
		message = GetStatusMsg(status)
	}

	return gin.H{
		"status":  status,
		"message": message,
		"data":    data,
	}
}
