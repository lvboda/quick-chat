package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/lvboda/quick-chat/utils"
	"github.com/lvboda/quick-chat/utils/status"
)

func jwtMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenList := strings.Split(c.Request.Header.Get("Authorization"), " ")

		// 跳过注册和登录
		if strings.Contains(c.Request.URL.Path, "/user/register") || strings.Contains(c.Request.URL.Path, "/user/login") || strings.Contains(c.Request.URL.Path, "/chat") {
			c.Next()
			return
		}

		// 跳过ws和静态资源
		if strings.Contains(c.Request.URL.Path, "/chat") || strings.Contains(c.Request.URL.Path, "/assets") {
			c.Next()
			return
		}

		// 判断token格式
		if len(tokenList) == 0 || len(tokenList) != 2 || tokenList[0] != "Bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, status.GetResponse(status.ERROR_TOKEN_TYPE_WRONG, nil, nil))
			return
		}

		// 过期或其他错误
		if claims, code := utils.ParseToken(tokenList[1]); code != status.SUCCESS {
			c.AbortWithStatusJSON(http.StatusUnauthorized, status.GetResponse(code, nil, nil))
			return
		} else {
			// next
			c.Set("claims", claims)
			c.Next()
		}
	}
}
