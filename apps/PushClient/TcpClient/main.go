package main

import (
	"io"
	"log"
	"net"
	"time"
)

func main(){

	//用于重连
Reconnection:

	host := "192.168.0.9:1243"
	hawkServer, err := net.ResolveTCPAddr("tcp", host)
	if err != nil {
		log.Printf("hawk server [%s] resolve error: [%s]", host, err.Error())
		time.Sleep(1 * time.Second)
		goto Reconnection
	}

	//连接服务器
	connection, err := net.DialTCP("tcp", nil, hawkServer)
	if err != nil {
		log.Printf("connect to hawk server error: [%s]", err.Error())
		time.Sleep(1 * time.Second)
		goto Reconnection
	}
	log.Println("[连接成功] 连接服务器成功")

	//创建客户端实例
	client := &TcpClient{
		Connection: connection,
		HawkServer: hawkServer,
		StopChan:   make(chan struct{}),
		CmdChan: make(chan string),
	}

	//启动接收
	go func(conn *TcpClient){
		for{
			recv := make([]byte, 1024)
			for {
				n, err := conn.Connection.Read(recv)
				if err != nil{
					if err == io.EOF {
						log.Println(conn.Addr(), " 断开了连接!")
						conn.Close()
						return
					}
				}
				if n > 0 && n < 1025 {
					conn.CmdChan <- string(recv[:n])
				}
			}
		}
	}(client)

	// 发送心跳
	go func(conn *TcpClient){
		i := 0
		heartBeatTick := time.Tick(10 * time.Second)
		for {
			select {
			case <-heartBeatTick:
				if _, err := conn.Send([]byte("beat")); err != nil {
					RConn <- true
					return
				}
				i++
			case <-conn.StopChan:
				return
			}
		}
	}(client)

	for {
		select {
		case a := <- RConn:
			log.Println("global.RConn = ", a)
			goto Reconnection
		}
	}

	//等待退出
	<-client.StopChan

}

// 重连
var RConn = make(chan bool)

type TcpClient struct {
	Connection *net.TCPConn
	HawkServer *net.TCPAddr
	StopChan   chan struct{}
	CmdChan chan string
	Token string
}

func (c *TcpClient) Send(b []byte) (int, error) {
	return c.Connection.Write(b)
}

func (c *TcpClient) Read(b []byte) (int, error) {
	return c.Connection.Read(b)
}

func (c *TcpClient) Addr() string {
	return c.Connection.RemoteAddr().String()
}

func (c *TcpClient) Close(){
	c.Connection.Close()
	RConn <- true
}
