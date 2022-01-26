package model

import (
	"fmt"
	"net"
)

type UdpClient struct {
	Conn *net.UDPAddr
	UserID int64 // 用户id唯一的
	IP string // 当前连接的ip
	DeviceID string // 当前连接的设备id
	Source string // 设备类型
	UDPListener *net.UDPConn
}

func (udp *UdpClient) GetIP() string {
	return udp.IP
}

func (udp *UdpClient) Send(msg []byte) {
	_,_=udp.UDPListener.WriteToUDP(msg, udp.Conn)
}

func (udp *UdpClient) GetConn() *net.UDPAddr {
	return udp.Conn
}

// 使用哈希表存储
type udpNode struct {
	Data map[int64]*UdpClient // UserID 用户id作为key
}

type UdpHashTable struct {
	Table map[int64]*udpNode // UserID 用户id 哈希后作为key
	Size  int64
}

func (table *UdpHashTable) hashFunction(uid int64) int64 {
	return uid % table.Size
}

func (table *UdpHashTable) Insert(value *UdpClient){
	h := table.hashFunction(value.UserID)
	element, ok := table.Table[h]
	if !ok {
		element = &udpNode{
			Data: make(map[int64]*UdpClient),
		}
		table.Table[h] = element
	}
	element.Data[value.UserID] = value
}

func (table *UdpHashTable) Get(uid int64) (date *UdpClient, err error){
	if t, ok := table.Table[table.hashFunction(uid)]; ok {
		if client, ok := t.Data[uid]; ok {
			return client, nil
		}
	}
	return nil, fmt.Errorf("not fond")
}

func (table *UdpHashTable) Del(value *UdpClient) {
	h := table.hashFunction(value.UserID)
	element, ok := table.Table[h]
	if ok {
		delete(element.Data, value.UserID)
	}
}

func InitUdpConnTable() *UdpHashTable {
	table := make(map[int64]*udpNode, HashSize)
	return &UdpHashTable{Table: table, Size: int64(HashSize)}
}

func udpTraverse(hash *TcpHashTable) {
	for k := range hash.Table {
		if hash.Table[k] != nil {
			for k,v := range hash.Table[k].Data {
				fmt.Printf("%v (%v) -> ", k, v)
			}
		}
		fmt.Println()
	}
}
