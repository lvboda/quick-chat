package status

import (
	"github.com/gin-gonic/gin"
)

const (
	// ws处理类型
	WS_PROCESS_HEART      = 0
	WS_PROCESS_SINGLE_MSG = 1
	WS_PROCESS_GROUP_MSG  = 2
	WS_PROCESS_CLOSE      = 3

	// 通用code
	SUCCESS = 200
	ERROR   = 500

	// code = 5000... 通用错误
	ERROR_REQUEST_PARAM = 5000
	ERROR_FILE_PARSE    = 5001
	ERROR_FILE_UPLOAD   = 5002

	// code= 1000... 用户模块错误
	ERROR_USERNAME_USED    = 1000
	ERROR_PASSWORD_WRONG   = 1001
	ERROR_USER_NOT_EXIST   = 1002
	ERROR_TOKEN_CREATE     = 1003
	ERROR_TOKEN_RUNTIME    = 1004
	ERROR_TOKEN_PARSE      = 1005
	ERROR_TOKEN_WRONG      = 1006
	ERROR_TOKEN_TYPE_WRONG = 1007
	ERROR_USER_NO_RIGHT    = 1008
	ERROR_USER_UPDATE      = 1009
	ERROR_USER_DELETE      = 1010
	ERROR_USER_REGISTER    = 1011

	// code = 2000... 关系模块错误
	ERROR_RELATION_ADD             = 2000
	ERROR_RELATION_DELETE          = 2001
	ERROR_RELATION_SELECT          = 2002
	ERROR_RELATION_VALIDATE_SELECT = 2003
	ERROR_RELATION_VALIDATE_SEND   = 2004

	// code = 3000... 聊天模块错误
	ERROR_CHAT_WS = 3000
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
	ERROR_TOKEN_PARSE:      "TOKEN解析失败",
	ERROR_TOKEN_RUNTIME:    "TOKEN已过期 请重新登陆",
	ERROR_TOKEN_WRONG:      "TOKEN不正确 请重新登陆",
	ERROR_TOKEN_TYPE_WRONG: "TOKEN格式错误 请重新登陆",
	ERROR_USER_NO_RIGHT:    "该用户无权限",
	ERROR_USER_UPDATE:      "用户信息修改失败",
	ERROR_USER_DELETE:      "用户注销失败",
	ERROR_USER_REGISTER:    "用户注册失败",

	ERROR_RELATION_ADD:             "好友添加失败",
	ERROR_RELATION_DELETE:          "好友删除失败",
	ERROR_RELATION_SELECT:          "好友查询失败",
	ERROR_RELATION_VALIDATE_SELECT: "验证信息查询失败",
	ERROR_RELATION_VALIDATE_SEND:   "验证信息发送失败",

	ERROR_CHAT_WS: "创建连接发生错误",
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
