package model

import (
	"github.com/gorilla/websocket"
	"github.com/mangenotwork/extras/apps/Push/mq"
	"net"
)

var TopicMap map[string]*Topic

type Topic struct {
	Name string `json:"topic_name"`
	ID string `json:"topic_id"` // 是uuid
	WsClient map[string]*WsClient   // deviceId : *WsClient
	TcpClient map[string]*TcpClient  // deviceId :
	UdpClient map[string]*net.UDPAddr
}


func (t *Topic) Send(msg *mq.MQMsg) {
	go func() {
		for _,wsClient := range t.WsClient {
			wsClient.TopicSend(msg)
		}
	}()

	go func() {
		for _,tcpClient := range t.TcpClient {
			tcpClient.TopicSend(msg)
		}
	}()

	//for _,t.udpClient := range t.UdpClient {
	//
	//}
}

type TopicData struct {
	TopicName string `json:"topic_name"`
	Message string `json:"message"`
	SendTime string `json:"send_time"`
}

func (ws *WsClient) TopicSend(data *mq.MQMsg){
	if ws == nil {
		return
	}
	msg := CmdData{
		Cmd: "TopicMessage",
		Data: data,
	}
	_=ws.Conn.WriteMessage(websocket.BinaryMessage, msg.Byte())
}

func (ws *TcpClient) TopicSend(data *mq.MQMsg){
	if ws == nil {
		return
	}
	msg := CmdData{
		Cmd: "TopicMessage",
		Data: data,
	}
	_,_=ws.Conn.Write(msg.Byte())
}

