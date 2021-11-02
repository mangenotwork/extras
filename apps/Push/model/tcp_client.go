package model

import (
	"net"
)

var AllTcpClient = make(map[string]*TcpClient)

type TcpClient struct {
	Conn net.Conn
	IP string
}

func NewTcpClient() *TcpClient {
	return new(TcpClient)
}

func (tcp *TcpClient) AddConn(conn net.Conn) *TcpClient {
	tcp.Conn = conn
	return tcp
}

func (tcp *TcpClient) SetIP(ip string) *TcpClient {
	tcp.IP = ip
	return tcp
}

func (tcp *TcpClient) Send(msg CmdData) {
	_,_=tcp.Conn.Write(msg.Byte())
}

func (tcp *TcpClient) SendMessage(str string) {
	msg := CmdData{
		Cmd: "Message",
		Data: str,
	}
	_,_=tcp.Conn.Write(msg.Byte())
}

func (tcp *TcpClient) IntoAllClient(device string) {
	AllTcpClient[device] = tcp
}

func (tcp *TcpClient) GetWsConn() *WsClient {
	return nil
}

func (tcp *TcpClient) GetTcpConn() *TcpClient {
	return tcp
}