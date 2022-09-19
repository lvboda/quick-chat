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
		if c.Request.URL.Path == "/user/register" || c.Request.URL.Path == "/user/login" {
			c.Next()
			return
		}

		// 判断token格式
		if len(tokenList) == 0 || len(tokenList) != 2 || tokenList[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, status.GetResponse(status.ERROR_TOKEN_TYPE_WRONG, nil, nil))
			c.Abort()
			return
		}

		claims, errCode := utils.ParseToken(tokenList[1])

		// 过期或其他错误
		if errCode != status.SUCCESS {
			c.JSON(http.StatusUnauthorized, status.GetResponse(errCode, nil, nil))
			c.Abort()
			return
		}

		// next
		c.Set("claims", claims)
		c.Next()
	}
}
