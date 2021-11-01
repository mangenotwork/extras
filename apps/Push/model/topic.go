package model

import (
	"github.com/mangenotwork/extras/apps/Push/mq"
	"net"
)

var TopicMap map[string]*Topic

type Topic struct {
	Name string `json:"topic_name"`
	ID string `json:"topic_id"` // 是uuid
	WsClient map[string]*WsClient   // deviceId : *WsClient
	TcpClient map[string]*net.Conn  // deviceId :
	UdpClient map[string]*net.UDPAddr
}


func (t *Topic) Send(msg *mq.MQMsg) {
	for _,wsClient := range t.WsClient {
		wsClient.TopicSend(msg)
	}
	//for _,t.tcpClient := range t.TcpClient {
	//
	//}
	//for _,t.udpClient := range t.UdpClient {
	//
	//}
}