package model

import (
	"encoding/json"
	"sync"

	"github.com/gorilla/websocket"
	"github.com/lvboda/quick-chat/utils"
	"gopkg.in/fatih/set.v0"
)

type Node struct {
	Conn      *websocket.Conn
	DataQueue chan []byte
	GroupSets set.Interface
}

func NewNode(conn *websocket.Conn) *Node {
	return &Node{
		Conn:      conn,
		DataQueue: make(chan []byte, 50),
		GroupSets: set.New(set.ThreadSafe),
	}
}

type Message struct {
	Id          string `json:"id"`
	SenderId    string `json:"senderId"`
	ReceiverId  string `json:"receiverId"`
	Content     string `json:"content"`
	Extra       string `json:"extra"`
	ContentType int    `json:"contentType"`
	ProcessType int    `json:"processType"`
	SendTime    string `json:"sendTime"`
	Resource    []byte `json:"resource"`
}

func ToMessage(data []byte) Message {
	var message Message
	err := json.Unmarshal(data, &message)
	if err != nil {
		utils.Logger.Errorln("json转Message结构体发生错误: ", err)
	}

	message.Resource = data
	return message
}

type OfflineGroup struct {
	OfflineMap map[string][]Message
	Locker     sync.RWMutex
}

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

func (og *OfflineGroup) Add(msg Message) {
	og.Locker.Lock()
	og.OfflineMap[msg.ReceiverId] = append(og.OfflineMap[msg.ReceiverId], msg)
	og.Locker.Unlock()
}

func (og *OfflineGroup) Delete(id string) {
	og.Locker.Lock()
	delete(og.OfflineMap, id)
	og.Locker.Unlock()
}

type NodeGroup struct {
	NodeMap map[string]*Node
	Locker  sync.RWMutex
}

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

func (ng *NodeGroup) Delete(id string) {
	ng.Locker.Lock()
	if _, ok := ng.NodeMap[id]; ok {
		ng.NodeMap[id].Conn.Close()
		delete(ng.NodeMap, id)
	}
	ng.Locker.Unlock()
}

func (ng *NodeGroup) SendMessage(msg Message) (ok bool) {
	ng.Locker.Lock()
	if node, has := ng.NodeMap[msg.ReceiverId]; has {
		node.DataQueue <- msg.Resource
		ok = true
	}
	ng.Locker.Unlock()
	return
}

func (ng *NodeGroup) SendGroupMessage(msg Message) {
	ng.Locker.Lock()
	for _, node := range ng.NodeMap {
		if node.GroupSets.Has(msg.ReceiverId) {
			node.DataQueue <- msg.Resource
		}
	}
	ng.Locker.Unlock()
}
