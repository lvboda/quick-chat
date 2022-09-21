package apiV1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/lvboda/quick-chat/service"
	"github.com/lvboda/quick-chat/utils"
)

func Chat(c *gin.Context) {
	conn, err := createConnect(c)
	if err != nil {
		utils.Logger.Errorln("创建websocket连接发生错误: ", err)
		return
	}

	service.Chat(c, conn)
}

// createConnect 创建websocket连接
func createConnect(c *gin.Context) (*websocket.Conn, error) {
	var upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin:     func(r *http.Request) bool { return true },
	}

	return upgrader.Upgrade(c.Writer, c.Request, nil)
}
