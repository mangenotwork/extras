package model

import "net"

var (
	AllUdpClient = make(map[string]*UdpClient)
	UDPListener *net.UDPConn
)


type UdpClient struct {
	Conn *net.UDPAddr
	IP string
}

func NewUdpClient() *UdpClient {
	return new(UdpClient)
}

func (udp *UdpClient) AddConn(conn *net.UDPAddr) *UdpClient {
	udp.Conn = conn
	return udp
}

func (udp *UdpClient) SetIP(ip string) *UdpClient {
	udp.IP = ip
	return udp
}

func (udp *UdpClient) Send(msg CmdData) {
	_,_=UDPListener.WriteToUDP(msg.Byte(), udp.Conn)
}

func (udp *UdpClient) SendMessage(str string) {
	msg := CmdData{
		Cmd: "Message",
		Data: str,
	}
	_,_=UDPListener.WriteToUDP(msg.Byte(), udp.Conn)
}

func (udp *UdpClient) IntoAllClient(device string) {
	AllUdpClient[device] = udp
}

func (udp *UdpClient) GetWsConn() *WsClient {
	return nil
}

func (udp *UdpClient) GetTcpConn() *TcpClient {
	return nil
}

func (udp *UdpClient) GetUdpConn() *UdpClient {
	return udp
}

func (udp *UdpClient) Who() string {
	return "udp"
}