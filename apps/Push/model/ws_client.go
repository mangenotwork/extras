package model

import (
	"github.com/gorilla/websocket"
	"github.com/mangenotwork/extras/apps/Push/mq"
)

var AllWsClient = make(map[string]*WsClient)

type WsClient struct {
	Conn *websocket.Conn
	IP string
}

func (ws *WsClient) SendMessage(str string) {
	msg := CmdData{
		Cmd: "Message",
		Data: str,
	}
	_=ws.Conn.WriteMessage(websocket.BinaryMessage, msg.Byte())
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