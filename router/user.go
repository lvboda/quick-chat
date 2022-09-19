package router

import (
	"github.com/gin-gonic/gin"
	apiV1 "github.com/lvboda/quick-chat/api/v1"
)

// registerUserRoutes 注册用户模块路由
func registerUserRoutes(router *gin.RouterGroup) {
	router.GET("/user/:uid", apiV1.QueryUserByUid)
	router.PUT("/user/:id", apiV1.EditUserById)
	router.DELETE("/user/:id", apiV1.RemoveUserById)
	router.POST("/user/register", apiV1.Register)
	router.POST("/user/login", apiV1.Login)
}
