package router

import (
	"github.com/gin-gonic/gin"
	apiV1 "github.com/lvboda/quick-chat/api/v1"
)

// registerRelationRoutes 注册关系模块路由
func registerRelationRoutes(router *gin.RouterGroup) {
	router.POST("/relation/validate", apiV1.SendValidateInfo)
	router.POST("/relation/friend", apiV1.AddFriend)
	router.DELETE("/relation/friend", apiV1.RemoveFriend)
	router.GET("/relation/list/friend/:uid", apiV1.QueryFriendList)
	router.GET("/relation/list/validate/:uid", apiV1.QueryValidateInfoList)
}
