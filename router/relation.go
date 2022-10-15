package router

import (
	"github.com/gin-gonic/gin"
	apiV1 "github.com/lvboda/quick-chat/api/v1"
)

// registerRelationRoutes 注册关系模块路由
func registerRelationRoutes(router *gin.RouterGroup) {
	router.POST("/relation/validate", apiV1.SendValidate)
	router.POST("/relation", apiV1.AddRelation)
	router.DELETE("/relation", apiV1.RemoveRelation)
	router.POST("/relation/list", apiV1.QueryRelationList)
}
