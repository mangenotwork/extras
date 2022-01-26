package engine

import (
	"encoding/json"
	"github.com/mangenotwork/extras/apps/IM-Conn/handler"
	"github.com/mangenotwork/extras/apps/IM-Conn/model"
	"github.com/mangenotwork/extras/common/conf"
	"github.com/mangenotwork/extras/common/logger"
	"github.com/mangenotwork/extras/common/utils"
	"net"
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

		client := model.UdpClient{
			Conn : remoteAddr,
			IP : remoteAddr.String(),
			UDPListener : UDPListener,
		}

		logger.Info("<%s> %s\n", remoteAddr, data[:n])

		if n > 0 && n < 10241 {
			data := data[:n]
			logger.Info(string(data))
			cmdData := &model.CmdData{}
			jsonErr := json.Unmarshal(data, &cmdData)
			if jsonErr != nil {
				client.Send([]byte("非法数据格式"))
				continue
			}
			go handler.UdpHandler(client, cmdData)
		}else{
			client.Send([]byte("传入的数据太小或太大, 建议 1~10240个字节"))
		}

	}


}
