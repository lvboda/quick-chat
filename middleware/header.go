package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"
)

func headerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if strings.Contains(c.Request.URL.Path, "/assets") && strings.Contains(c.Request.URL.Path, ".html") {
			c.Writer.Header().Set("Cross-Origin-Opener-Policy", "same-origin")
			c.Writer.Header().Set("Cross-Origin-Embedder-Policy", "require-corp")
		}
	}
}
