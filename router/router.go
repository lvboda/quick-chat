package router

import (
	"github.com/gin-gonic/gin"
	"github.com/lvboda/quick-chat/utils"
)

// RegisterRoutes 注册路由
func RegisterRoutes(app *gin.Engine) {
	registerStaticRoutes(app)
	registerV1Routes(app)
}

// registerStaticRoutes 注册静态资源路由
func registerStaticRoutes(app *gin.Engine) {
	app.Static("/assets", utils.StaticAssetsPath)
}

// registerV1Routes 路由分组v1
func registerV1Routes(app *gin.Engine) {
	router := app.Group("/api/v1")
	registerUserRoutes(router)
	registerUploadRoutes(router)
	registerRelationRoutes(router)
	registerChatRoutes(router)
	registerCommunityRoutes(router)
}
