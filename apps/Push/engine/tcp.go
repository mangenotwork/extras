package engine

import (
	"encoding/json"
	"io"
	"net"

	"github.com/mangenotwork/extras/apps/Push/model"
	"github.com/mangenotwork/extras/apps/Push/service"
	"github.com/mangenotwork/extras/common/conf"
	"github.com/mangenotwork/extras/common/logger"
)

type TcpServer struct {
	Listener   *net.TCPListener
	HawkServer *net.TCPAddr
}

func StartTcpServer(){
	go func() {
		logger.Info("StartTcpServer")

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
		logger.Info("start TCP server successful.")

		//接收请求
		for {

			var (
				device *model.Device
				client model.Client
				deviceId string
			)

			//来自客户端的连接
			conn, err := tcpServer.Listener.Accept()
			if err != nil {
				logger.Error("[连接失败]:", err.Error())
				continue
			}
			logger.Info("[连接成功]: ", conn.RemoteAddr().String(), conn)

			wsClient := model.NewTcpClient().AddConn(conn).SetIP(conn.RemoteAddr().String())
			client = wsClient

			go func(){

				defer func() {
					if v := recover(); v != nil {
						logger.Error("捕获了一个异常：", v)
					}
					_=conn.Close()
				}()

				recv := make([]byte, 1024*10)
				for {
					n, err := conn.Read(recv)
					logger.Info(n, err)
					if err != nil{
						if err == io.EOF {
							logger.Info(conn.RemoteAddr().String(), " 断开了连接!")
							// 如果认证了设备则清理设备
							if device != nil {
								logger.Info("释放客户端连接")
								device.OffLine() // 下线记录
								delete(model.AllWsClient, deviceId)
								device.Discharge("tcp") // 连接离开topic,group
							}
							_=conn.Close()
							return
						}
					}
					if n > 0 && n < 10241 {
						data := recv[:n]
						logger.Info(string(data))
						cmdData := &model.CmdData{}
						jsonErr := json.Unmarshal(data, &cmdData)
						if jsonErr != nil {
							_,_=conn.Write(model.CmdDataMsg("非法数据格式"))
							continue
						}
						device = service.Interactive(cmdData, client)

					}else{
						_,_=conn.Write(model.CmdDataMsg("传入的数据太小或太大, 建议 1~10240个字节"))

					}
				}
			}()
		}
	}()
}

