package router

import (
	"github.com/gin-gonic/gin"
	apiV1 "github.com/lvboda/quick-chat/api/v1"
)

// registerCommunityRoutes 注册群聊模块路由
func registerCommunityRoutes(router *gin.RouterGroup) {
	router.GET("/community/:cid", apiV1.QueryCommunityByCid)
	router.POST("/community", apiV1.CreateCommunity)
	router.PUT("/community/:cid", apiV1.EditCommunityById)
	router.DELETE("/community/:cid", apiV1.RemoveCommunityById)
}
