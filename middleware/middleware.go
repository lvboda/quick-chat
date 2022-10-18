package middleware

import (
	"github.com/gin-gonic/gin"
)

func RegisterMiddleware(app *gin.Engine) {
	app.Use(corsMiddleware())
	app.Use(jwtMiddleware())
	app.Use(headerMiddleware())
}
