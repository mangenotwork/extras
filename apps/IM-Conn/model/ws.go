package model

import (
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/mangenotwork/extras/common/utils"
)

type WsClient struct {
	Conn *websocket.Conn
	UserID string // 用户id唯一的
	IP string // 当前连接的ip
	DeviceID string // 当前连接的设备id
	DeviceType string // 设备类型
	HeartBeat chan []byte // 心跳包
}

func (ws *WsClient) GetIP() string {
	return ws.IP
}

func (ws *WsClient) Send(msg []byte) {
	_=ws.Conn.WriteMessage(websocket.BinaryMessage, msg)
}

func (ws *WsClient) GetConn() *websocket.Conn {
	return ws.Conn
}

// 使用哈希表存储
type wsNode struct {
	Data map[string]*WsClient // UserID 用户id作为key
}

type WsHashTable struct {
	Table map[int64]*wsNode // UserID 用户id 哈希后作为key
	Size  int64
}

// 哈希算法 userid余表Size
func (table *WsHashTable) hashFunction(uid string) int64 {
	uidInt := utils.Str2Int64(uid)
	return uidInt % table.Size
}

func (table *WsHashTable) Insert(value *WsClient){
	h := table.hashFunction(value.UserID)
	element, ok := table.Table[h]
	if !ok {
		element = &wsNode{
			Data: make(map[string]*WsClient),
		}
		table.Table[h] = element
	}
	element.Data[value.UserID] = value
}

func (table *WsHashTable) Get(uid string) (date *WsClient, err error){
	if t, ok := table.Table[table.hashFunction(uid)]; ok {
		if client, ok := t.Data[uid]; ok {
			return client, nil
		}
	}
	return nil, fmt.Errorf("not fond")
}

func (table *WsHashTable) Del(value *WsClient) {
	h := table.hashFunction(value.UserID)
	element, ok := table.Table[h]
	if ok {
		delete(element.Data, value.UserID)
	}
}

func InitWsConnTable() *WsHashTable {
	table := make(map[int64]*wsNode, HashSize)
	return &WsHashTable{Table: table, Size: int64(HashSize)}
}

func wsTraverse(hash *WsHashTable) {
	for k := range hash.Table {
		if hash.Table[k] != nil {
			for k,v := range hash.Table[k].Data {
				fmt.Printf("%v (%v) -> ", k, v)
			}
		}
		fmt.Println()
	}
}