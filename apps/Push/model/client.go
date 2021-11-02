package model

type Client interface {
	SendMessage(str string)
	IntoAllClient(device string)
	GetWsConn() *WsClient
	GetTcpConn() *TcpClient
	GetUdpConn() *UdpClient
	Who() string
}

func NewClient(clientType string) Client {
	switch clientType {
	case "ws":
		return &WsClient{}
	case "tcp":
		return &TcpClient{}
	case "udp":
		return &UdpClient{}
	}
	return nil
}
