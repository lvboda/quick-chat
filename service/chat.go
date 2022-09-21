package service

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/lvboda/quick-chat/model"
	"github.com/lvboda/quick-chat/utils"
	"github.com/lvboda/quick-chat/utils/status"
)

var globalNodeGroup = model.NewNodeGroup()
var offlineGroup = model.NewOfflineGroup()
var closeChan = make(chan bool)

func Chat(c *gin.Context, conn *websocket.Conn) {
	uid := c.Param("uid")
	// if ok := utils.CheckAuthByUserId(c, uid); !ok {
	// 	return
	// }

	node, ok := globalNodeGroup.Add(uid, conn)
	if !ok {
		return
	}

	go sendLoop(uid, node)
	go receiveLoop(uid, node)

	pushOfflineMsg(uid)
	wait(uid)
}

func pushOfflineMsg(uid string) {
	if msgQueue, has := offlineGroup.OfflineMap[uid]; has {
		for _, msg := range msgQueue {
			sendMessage(msg)
		}

		offlineGroup.Delete(uid)
	}
}

func sendLoop(uid string, node *model.Node) {
	for {
		data := <-node.DataQueue
		err := node.Conn.WriteMessage(websocket.TextMessage, data)

		if err != nil {
			closeChan <- true
			return
		}
	}
}

func receiveLoop(uid string, node *model.Node) {
	for {
		_, data, err := node.Conn.ReadMessage()
		if err != nil {
			closeChan <- true
			return
		}

		dispatch(data)
	}
}

func dispatch(data []byte) {
	msg := model.ToMessage(data)

	switch msg.ProcessType {
	case status.WS_PROCESS_SINGLE_MSG:
		sendMessage(msg)
	case status.WS_PROCESS_GROUP_MSG:
		sendGroupMessage(msg)
	case status.WS_PROCESS_CLOSE:
		closeChan <- true
	case status.WS_PROCESS_HEART:
		// ❤️
	}
}

func sendMessage(msg model.Message) {
	if ok := globalNodeGroup.SendMessage(msg); !ok {
		offlineGroup.Add(msg)
	}
}

func sendGroupMessage(msg model.Message) {
	globalNodeGroup.SendGroupMessage(msg)
}

func wait(uid string) {
	for {
		if <-closeChan {
			utils.Logger.Debugln(111999)
			globalNodeGroup.Delete(uid)
			return
		}
	}
}
