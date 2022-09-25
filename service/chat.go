package service

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/lvboda/quick-chat/model"
	"github.com/lvboda/quick-chat/utils"
	"github.com/lvboda/quick-chat/utils/status"
)

var globalNodeGroup = model.NewNodeGroup()
var globalOfflineGroup = model.NewOfflineGroup()
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

	utils.Logger.Infof("ws:用户%s连接成功", uid)

	go sendLoop(uid, node)
	go receiveLoop(uid, node)

	pushGroupId(uid, node)
	pushOfflineMsg(uid)
	wait(uid)
}

// pushGroupId 添加群id
func pushGroupId(uid string, node *model.Node) {
	var query struct {
		FriendId     string
		RelationType int
		RoleType     int
	}
	query.FriendId = uid
	query.RelationType = 2
	query.RoleType = 2
	relationList, _ := model.RelationEntity{}.SelectListBy(query, query.RoleType)

	for _, relation := range relationList {
		if relation.CommunityInfo.CommunityId != "" {
			node.GroupSets.Add(relation.CommunityInfo.CommunityId)
		}
	}
}

// pushOfflineMsg 离线消息推送
func pushOfflineMsg(uid string) {
	if msgQueue, has := globalOfflineGroup.OfflineMap[uid]; has {
		for _, msg := range msgQueue {
			sendMessage(msg)
		}

		globalOfflineGroup.Delete(uid)
	}
}

// sendLoop 发送线程
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

// receiveLoop 接收线程
func receiveLoop(uid string, node *model.Node) {
	for {
		_, data, err := node.Conn.ReadMessage()
		if err != nil {
			closeChan <- true
			return
		}

		dispatchProcess(data)
	}
}

// dispatchProcess 分发不同的处理函数
func dispatchProcess(data []byte) {
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

// sendMessage 发送消息
func sendMessage(msg model.Message) {
	if ok := globalNodeGroup.SendMessage(msg); !ok {
		globalOfflineGroup.Add(msg)
	}
}

// sendGroupMessage 发送群消息
func sendGroupMessage(msg model.Message) {
	globalNodeGroup.SendGroupMessage(msg)
}

// wait 阻塞主线程 直到close
func wait(uid string) {
	for {
		if <-closeChan {
			globalNodeGroup.Delete(uid)
			utils.Logger.Infof("ws:用户%s断开连接", uid)
			return
		}
	}
}
