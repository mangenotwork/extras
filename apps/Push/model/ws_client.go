package model

import "github.com/gorilla/websocket"

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

func (ws *WsClient) TopicSend(data string){
	msg := CmdData{
		Cmd: "TopicMessage",
		Data: data,
	}
	_=ws.Conn.WriteMessage(websocket.BinaryMessage, msg.Byte())
}