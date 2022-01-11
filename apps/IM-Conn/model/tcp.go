package model

import (
	"fmt"
	"net"
)

type TcpClient struct {
	conn net.Conn
	UserID int64 // 用户id唯一的
	IP string // 当前连接的ip
	DeviceID string // 当前连接的设备id
	DeviceType string // 设备类型
}

func (tcp *TcpClient) GetIP() string {
	return tcp.IP
}

func (tcp *TcpClient) Send(msg []byte) {
	_,_=tcp.conn.Write(msg)
}

func (tcp *TcpClient) Conn() net.Conn {
	return tcp.conn
}

// 使用哈希表存储
type tcpNode struct {
	Data map[int64]*TcpClient // UserID 用户id作为key
}

type TcpHashTable struct {
	Table map[int64]*tcpNode // UserID 用户id 哈希后作为key
	Size  int64
}

func (table *TcpHashTable) hashFunction(uid int64) int64 {
	return uid % table.Size
}

func (table *TcpHashTable) Insert(value *TcpClient){
	h := table.hashFunction(value.UserID)
	element, ok := table.Table[h]
	if !ok {
		element = &tcpNode{
			Data: make(map[int64]*TcpClient),
		}
		table.Table[h] = element
	}
	element.Data[value.UserID] = value
}

func (table *TcpHashTable) Get(uid int64) (date *TcpClient, err error){
	if t, ok := table.Table[table.hashFunction(uid)]; ok {
		if client, ok := t.Data[uid]; ok {
			return client, nil
		}
	}
	return nil, fmt.Errorf("not fond")
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
