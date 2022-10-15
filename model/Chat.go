package model

import (
	"encoding/json"
	"sync"

	"github.com/gorilla/websocket"
	"github.com/lvboda/quick-chat/utils"
	"gopkg.in/fatih/set.v0"
)

// 一个用户对应一个node
type Node struct {
	// websocket连接
	Conn *websocket.Conn
	// 数据存储队列
	DataQueue chan []byte
	// 群id的set
	GroupSets set.Interface
}

// NewNode 创建新的node
func NewNode(conn *websocket.Conn) *Node {
	return &Node{
		Conn:      conn,
		DataQueue: make(chan []byte, 50),
		GroupSets: set.New(set.ThreadSafe),
	}
}

// 传输消息结构
type Message struct {
	// 消息id
	Id string `json:"id"`
	// 发送者id
	SenderId string `json:"senderId"`
	// 接受者id
	ReceiverId string `json:"receiverId"`
	// 消息内容
	Content string `json:"content"`
	// 附加信息
	Extra string `json:"extra"`
	// 消息类型 前端用来判断 后端不处理
	ContentType int `json:"contentType"`
	// 处理类型 见status
	ProcessType int `json:"processType"`
	// 发送时间
	SendTime string `json:"sendTime"`
	// 源数据
	Resource []byte `json:"resource"`
}

// ToMessage []byte转message
func ToMessage(data []byte) Message {
	var message Message
	err := json.Unmarshal(data, &message)
	if err != nil && string(data) != "heart" {
		utils.Logger.Errorln("ws:json转Message结构体发生错误: ", err)
	}

	message.Resource = data
	return message
}

// 离线消息存储
type OfflineGroup struct {
	OfflineMap map[string][]Message
	Locker     sync.RWMutex
}

// NewOfflineGroup 新建OfflineGroup 全局维护一个
func NewOfflineGroup() *OfflineGroup {
	var flag bool
	var og *OfflineGroup
	return func() *OfflineGroup {
		if flag {
			return og
		} else {
			flag = true
			og = &OfflineGroup{
				OfflineMap: map[string][]Message{},
				Locker:     sync.RWMutex{},
			}
			return og
		}
	}()
}

// Add 添加消息
func (og *OfflineGroup) Add(msg Message) {
	og.Locker.Lock()
	og.OfflineMap[msg.ReceiverId] = append(og.OfflineMap[msg.ReceiverId], msg)
	og.Locker.Unlock()
}

// Delete 删除消息
func (og *OfflineGroup) Delete(id string) {
	og.Locker.Lock()
	delete(og.OfflineMap, id)
	og.Locker.Unlock()
}

// 全局node存储 {key:uid value:node}
type NodeGroup struct {
	NodeMap map[string]*Node
	Locker  sync.RWMutex
}

// NewNodeGroup 新建NodeGroup 全局维护一个
func NewNodeGroup() *NodeGroup {
	var flag bool
	var ng *NodeGroup
	return func() *NodeGroup {
		if flag {
			return ng
		} else {
			flag = true
			ng = &NodeGroup{
				NodeMap: make(map[string]*Node),
				Locker:  sync.RWMutex{},
			}
			return ng
		}
	}()
}

// Add 添加node
func (ng *NodeGroup) Add(id string, conn *websocket.Conn) (node *Node, ok bool) {
	ng.Locker.Lock()
	if _, has := ng.NodeMap[id]; !has {
		node = NewNode(conn)
		ng.NodeMap[id] = node
		ok = true
	}
	ng.Locker.Unlock()
	return
}

// Delete 删除node
func (ng *NodeGroup) Delete(id string) {
	ng.Locker.Lock()
	if _, ok := ng.NodeMap[id]; ok {
		ng.NodeMap[id].Conn.Close()
		delete(ng.NodeMap, id)
	}
	ng.Locker.Unlock()
}

// SendMessage 发送消息
func (ng *NodeGroup) SendMessage(msg Message) (ok bool) {
	ng.Locker.Lock()
	if node, has := ng.NodeMap[msg.ReceiverId]; has {
		node.DataQueue <- msg.Resource
		ok = true
	}
	ng.Locker.Unlock()
	return
}

// SendGroupMessage 发送群消息
func (ng *NodeGroup) SendGroupMessage(msg Message) {
	ng.Locker.Lock()
	for _, node := range ng.NodeMap {
		if node.GroupSets.Has(msg.ReceiverId) {
			node.DataQueue <- msg.Resource
		}
	}
	ng.Locker.Unlock()
}
