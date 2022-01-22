package engine

import (
	"net"
	"time"

	"github.com/mangenotwork/extras/apps/IM-Conn/model"
	"github.com/mangenotwork/extras/common/conf"
	"github.com/mangenotwork/extras/common/logger"
	"github.com/mangenotwork/extras/common/utils"
)

func StartUDP(){
	go func() {
		logger.Info("StartUDP")
		RunUDPServer()
	}()
}

func RunUDPServer() {

	// 监听
	UDPListener, err := net.ListenUDP("udp", &net.UDPAddr{
		IP: net.ParseIP("0.0.0.0"),
		Port: utils.Str2Int(conf.Arg.UdpServer.Prod),
	})
	if err != nil {
		logger.Error(err)
		return
	}
	logger.Info("Local: <%s> \n", UDPListener.LocalAddr().String())

	// 读取数据
	data := make([]byte, 10240)
	for {

		n, remoteAddr, err := UDPListener.ReadFromUDP(data)
		if err != nil {
			logger.Error("error during read: %s", err)
		}

		client := &model.UdpClient{
			Conn : remoteAddr,
			IP : remoteAddr.String(),
			UDPListener : UDPListener,
		}


		logger.Info("<%s> %s\n", remoteAddr, data[:n])
		//_, err = UDPListener.WriteToUDP([]byte("world"), remoteAddr)
		//if err != nil {
		//	log.Printf(err.Error())
		//}


		if n > 0 && n < 10241 {
			data := data[:n]
			logger.Info(string(data))
			//cmdData := &model.CmdData{}
			//jsonErr := json.Unmarshal(data, &cmdData)
			//if jsonErr != nil {
			//	UDPSend(remoteAddr, model.CmdDataMsg("非法数据格式"))
			//	continue
			//}
			////device = service.Interactive(cmdData, client)
			//service.Interactive(cmdData, client)

			// TODO 获取来自客服端的身份信息,并验证
			model.UdpClientTable().Insert(client)

		}else{
			UDPListener.WriteToUDP([]byte("传入的数据太小或太大, 建议 1~10240个字节"), remoteAddr)

		}

		// TODO 处理心跳
		go func() {
			for {
				timer := time.NewTimer(10 * time.Second)

				select {
				case <-timer.C:
					// TODO 删除客户端对象
					model.UdpClientTable().Del(client)

					// 接收心跳
					//case <-rafter.heartBeat:
					//	// 重置
					//	log.Println("重置 = ", 10)
					//	timer.Reset(10 * time.Second)
				}
			}
		}()

	}


}
