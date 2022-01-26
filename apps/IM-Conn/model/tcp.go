package model

import (
	"fmt"
	"github.com/mangenotwork/extras/common/utils"
	"net"
)

type TcpClient struct {
	Conn net.Conn
	UserID string // 用户id唯一的
	IP string // 当前连接的ip
	DeviceID string // 当前连接的设备id
	Source string // 设备类型
	HeartBeat chan []byte // 心跳包
}

func (tcp *TcpClient) GetIP() string {
	return tcp.IP
}

func (tcp *TcpClient) Send(msg []byte) {
	_,_=tcp.Conn.Write(msg)
}

func (tcp *TcpClient) GetConn() net.Conn {
	return tcp.Conn
}

// 使用哈希表存储
type tcpNode struct {
	Data map[string]*TcpClient // UserID 用户id作为key
}

type TcpHashTable struct {
	Table map[int64]*tcpNode // UserID 用户id 哈希后作为key
	Size  int64
}

func (table *TcpHashTable) hashFunction(uid string) int64 {
	uidInt := utils.Str2Int64(uid)
	return uidInt % table.Size
}

func (table *TcpHashTable) Insert(value *TcpClient){
	h := table.hashFunction(value.UserID)
	element, ok := table.Table[h]
	if !ok {
		element = &tcpNode{
			Data: make(map[string]*TcpClient),
		}
		table.Table[h] = element
	}
	element.Data[value.UserID] = value
}

func (table *TcpHashTable) Get(uid string) (date *TcpClient, err error){
	if t, ok := table.Table[table.hashFunction(uid)]; ok {
		if client, ok := t.Data[uid]; ok {
			return client, nil
		}
	}
	return nil, fmt.Errorf("not fond")
}

func (table *TcpHashTable) Del(value *TcpClient) {
	h := table.hashFunction(value.UserID)
	element, ok := table.Table[h]
	if ok {
		delete(element.Data, value.UserID)
	}
}


func InitTcpConnTable() *TcpHashTable {
	table := make(map[int64]*tcpNode, HashSize)
	return &TcpHashTable{Table: table, Size: int64(HashSize)}
}

func tcpTraverse(hash *TcpHashTable) {
	for k := range hash.Table {
		if hash.Table[k] != nil {
			for k,v := range hash.Table[k].Data {
				fmt.Printf("%v (%v) -> ", k, v)
			}
		}
		fmt.Println()
	}
}
