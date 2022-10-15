package router

import (
	"github.com/gin-gonic/gin"
	apiV1 "github.com/lvboda/quick-chat/api/v1"
)

// registerUploadRoutes 注册上传模块路由
func registerUploadRoutes(router *gin.RouterGroup) {
	router.POST("/upload", apiV1.UploadFile)
	router.POST("/upload/temp", apiV1.UploadTempFile)
}
