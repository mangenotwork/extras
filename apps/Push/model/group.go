package model

import "net"

var GroupMap = make(map[string]*Group)

type Group struct {
	Name string `json:"group_name"`
	ID string `json:"group_id"` // 是uuid
	WsClient map[string]*WsClient   // deviceId : *WsClient
	TcpClient map[string]*TcpClient  // deviceId :
	UdpClient map[string]*net.UDPAddr
}