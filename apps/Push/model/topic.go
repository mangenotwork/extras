package model

import (
	"net"
)

var TopicMap = make(map[string]*Topic)

type Topic struct {
	Name string `json:"topic_name"`
	ID string `json:"topic_id"` // 是uuid
	WsClient map[string]*WsClient   // deviceId : *WsClient
	TcpClient map[string]*net.Conn  // deviceId :
	UdpClient map[string]*net.UDPAddr
}


