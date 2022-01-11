package model

import (
	"fmt"
	"github.com/gorilla/websocket"
	"sync"
)

type WsClient struct {
	conn *websocket.Conn
	UserID int64 // 用户id唯一的
	IP string // 当前连接的ip
	DeviceID string // 当前连接的设备id
	DeviceType string // 设备类型
}

func (ws *WsClient) GetIP() string {
	return ws.IP
}

func (ws *WsClient) Send(msg []byte) {
	_=ws.conn.WriteMessage(websocket.BinaryMessage, msg)
}

func (ws *WsClient) Conn() *websocket.Conn {
	return ws.conn
}

// 使用哈希表存储
type wsNode struct {
	Data map[int64]*WsClient // UserID 用户id作为key
}

type WsHashTable struct {
	Table map[int64]*wsNode // UserID 用户id 哈希后作为key
	Size  int64
}

func (table *WsHashTable) hashFunction(uid int64) int64 {
	return uid % table.Size
}

func (table *WsHashTable) Insert(value *WsClient){
	h := table.hashFunction(value.UserID)
	element, ok := table.Table[h]
	if !ok {
		element = &wsNode{
			Data: make(map[int64]*WsClient),
		}
		table.Table[h] = element
	}
	element.Data[value.UserID] = value
}


func InitWsConnTable() *WsHashTable {
	table := make(map[int64]*wsNode, 10)
	return &WsHashTable{Table: table, Size: 10}
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

func (table *WsHashTable) Get(uid int64) (date *WsClient, err error){
	if t, ok := table.Table[table.hashFunction(uid)]; ok {
		if client, ok := t.Data[uid]; ok {
			return client, nil
		}
	}
	return nil, fmt.Errorf("not fond")
}

var _wsConnTable = InitWsConnTable()
var _wsConnTableOnce sync.Once

func WsConnTable() *WsHashTable{
	_wsConnTableOnce.Do(func() {
		_wsConnTable = InitWsConnTable()
	})
	return _wsConnTable
}


