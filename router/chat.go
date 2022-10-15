package router

import (
	"github.com/gin-gonic/gin"
	apiV1 "github.com/lvboda/quick-chat/api/v1"
)

// registerWebsocketRouter 注册chat模块路由
func registerChatRoutes(router *gin.RouterGroup) {
	router.GET("/chat/:uid", apiV1.Chat)
}
