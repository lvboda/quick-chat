package router

import (
	"github.com/gin-gonic/gin"
)

// RegisterRoutes 注册路由
func RegisterRoutes(app *gin.Engine) {
	registerV1Routes(app)
}

// registerV1Routes 路由分组v1
func registerV1Routes(app *gin.Engine) {
	router := app.Group("/api/v1")
	registerUserRoutes(router)
	registerUploadRoutes(router)
}
