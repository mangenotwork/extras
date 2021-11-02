package model

type Client interface {
	SendMessage(str string)
	IntoAllClient(device string)
	GetWsConn() *WsClient
	GetTcpConn() *TcpClient
}

func NewClient(clientType string) Client {
	switch clientType {
	case "ws":
		return &WsClient{}
	case "tcp":
		return &TcpClient{}
	}
	return nil
}
