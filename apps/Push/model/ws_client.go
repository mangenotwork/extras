package model

import (
	"github.com/gorilla/websocket"
)

var AllWsClient = make(map[string]*WsClient)

type WsClient struct {
	Conn *websocket.Conn
	IP string
}

func NewWsClient() *WsClient{
	return new(WsClient)
}

func (ws *WsClient) AddConn(conn *websocket.Conn) *WsClient {
	ws.Conn = conn
	return ws
}

func (ws *WsClient) SetIP(ip string) *WsClient {
	ws.IP = ip
	return ws
}

func (ws *WsClient) Send(msg CmdData) {
	_=ws.Conn.WriteMessage(websocket.BinaryMessage, msg.Byte())
}

func (ws *WsClient) SendMessage(str string) {
	msg := CmdData{
		Cmd: "Message",
		Data: str,
	}
	_=ws.Conn.WriteMessage(websocket.BinaryMessage, msg.Byte())
}

func (ws *WsClient) IntoAllClient(device string) {
	AllWsClient[device] = ws
}

func (ws *WsClient) GetWsConn() *WsClient {
	return ws
}

func (ws *WsClient) GetTcpConn() *TcpClient {
	return nil
}

func (ws *WsClient) GetUdpConn() *UdpClient {
	return nil
}

func (ws *WsClient) Who() string {
	return "ws"
}