package engine

import (
	"github.com/mangenotwork/extras/common/conf"
	"log"
	"net"
)

type TcpServer struct {
	Listener   *net.TCPListener
	HawkServer *net.TCPAddr
}

func StartTcpServer(){
	go func() {
		log.Println("StartTcpServer")

		//类似于初始化套接字，绑定端口
		hawkServer, err := net.ResolveTCPAddr("tcp", "0.0.0.0:"+conf.Arg.TcpServer.Prod)
		if err != nil {
			panic("[终止] 出现致命错误: "+ err.Error())
		}

		//侦听
		listen, err := net.ListenTCP("tcp", hawkServer)
		if err != nil {
			panic("[终止] 出现致命错误: "+ err.Error())
		}

		//关闭
		defer listen.Close()

		tcpServer := &TcpServer{
			Listener:   listen,
			HawkServer: hawkServer,
		}
		log.Println("start TCP server successful.")

		//接收请求
		for {

			//来自客户端的连接
			conn, err := tcpServer.Listener.Accept()
			if err != nil {
				log.Println("[连接失败]:", err.Error())
				continue
			}
			log.Println("[连接成功]: ", conn.RemoteAddr().String(), conn)
		}

	}()
}
