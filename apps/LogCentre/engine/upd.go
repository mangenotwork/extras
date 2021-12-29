package engine

import (
	"net"
	"os"
	"strings"
	"time"

	"github.com/mangenotwork/extras/common/conf"
	"github.com/mangenotwork/extras/common/logger"
	"github.com/mangenotwork/extras/common/utils"
)

func StartUdp(){
	go func() {

		logger.Info("StartUdpServer")
		var err error
		var udpListener *net.UDPConn

		// 监听
		udpListener, err = net.ListenUDP("udp", &net.UDPAddr{
			IP: net.ParseIP("0.0.0.0"),
			Port: utils.Str2Int(conf.Arg.UdpServer.Prod),
		})
		if err != nil {
			logger.Error(err)
			return
		}
		logger.Info("Local: <%s> \n", udpListener.LocalAddr().String())

		// 读取数据
		data := make([]byte, 1024*20)
		for {

			n, remoteAddr, err := udpListener.ReadFromUDP(data)
			if err != nil {
				logger.Error("error during read: %s", err)
			}

			if n > 0 && n < 1024*20 {
				dataStr := string(data[:n])
				ip := remoteAddr.IP.String()
				name := ip
				nameList := strings.Split(dataStr, "|")
				if len(nameList) > 0 {
					name = nameList[0]
				}
				writer, _ := os.OpenFile( "logs/"+name+time.Now().Format("-20060102")+".log",
					os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
				d := []byte(ip+"|")
				d = append(d, data[:n]...)
				writer.Write(d)

			}else{
				logger.Error("传入的数据太小或太大, 建议 1~10240个字节")
			}
		}

	}()
}