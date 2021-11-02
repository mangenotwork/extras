package engine

import (
	"encoding/json"
	"github.com/mangenotwork/extras/apps/Push/model"
	"github.com/mangenotwork/extras/apps/Push/service"
	"github.com/mangenotwork/extras/common/conf"
	"github.com/mangenotwork/extras/common/utils"
	"log"
	"net"
)

func UDPSend(addr *net.UDPAddr, b []byte){
	if model.UDPListener != nil {
		_,_=model.UDPListener.WriteToUDP(b, addr)
	}
}

func StartUdpServer(){
	go func() {
		log.Println("StartUdpServer")
		var err error

		// 监听
		model.UDPListener, err = net.ListenUDP("udp", &net.UDPAddr{
			IP: net.ParseIP("0.0.0.0"),
			Port: utils.Str2Int(conf.Arg.UdpServer.Prod),
		})
		if err != nil {
			log.Println(err)
			return
		}
		log.Printf("Local: <%s> \n", model.UDPListener.LocalAddr().String())

		// 读取数据
		data := make([]byte, 10240)
		for {

			var (
				//device *model.Device
				client model.Client
				//deviceId string
			)

			n, remoteAddr, err := model.UDPListener.ReadFromUDP(data)
			if err != nil {
				log.Printf("error during read: %s", err)
			}
			log.Printf("<%s> %s\n", remoteAddr, data[:n])
			//_, err = UDPListener.WriteToUDP([]byte("world"), remoteAddr)
			//if err != nil {
			//	log.Printf(err.Error())
			//}

			wsClient := model.NewUdpClient().AddConn(remoteAddr).SetIP(remoteAddr.String())
			client = wsClient

			if n > 0 && n < 10241 {
				data := data[:n]
				log.Println(string(data))
				cmdData := &model.CmdData{}
				jsonErr := json.Unmarshal(data, &cmdData)
				if jsonErr != nil {
					UDPSend(remoteAddr, model.CmdDataMsg("非法数据格式"))
					continue
				}
				//device = service.Interactive(cmdData, client)
				service.Interactive(cmdData, client)
			}else{
				UDPSend(remoteAddr, model.CmdDataMsg("传入的数据太小或太大, 建议 1~10240个字节"))
			}

		}

	}()
}