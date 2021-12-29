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
				dataType := data[0:1]
				ip := remoteAddr.IP.String()

				// 日志写入文件
				if string(dataType) == "1" {
					dataStr := string(data[1:n])
					name := ip
					nameList := strings.Split(dataStr, "|")
					if len(nameList) > 0 {
						name = nameList[0]
					}
					writer, err := os.OpenFile( "logs/"+name+time.Now().Format("-20060102")+".log",
						os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
					if err != nil {
						logger.Error("打开日志文件失败  = ", err)
					}
					d := []byte(ip+"|")
					d = append(d, data[1:n]...)
					writer.Write(d)
				}

				// 日志是http日志 记录到db
				if string(dataType) == "2" {
					logger.Info("http 日志 ", string(data[1:n]))
					dataList := strings.Split(string(data[1:n]), "#")
					reqIp := dataList[0]
					reqMethod := dataList[1]
					reqUrl := dataList[2]
					reqCode := dataList[3]
					reqTime := utils.Str2Float64(dataList[4])
					logger.Info("reqIp = ", reqIp, " | reqMethod = ", reqMethod, " | reqUrl = ", reqUrl, "| reqCode", reqCode, "| reqTime = ", reqTime)

				}

				// 日志是grpc日志 记录到db
				if string(dataType) == "3" {
					logger.Info("http 日志 ", string(data[1:n]))
					dataList := strings.Split(string(data[1:n]), "#")
					state := dataList[0]
					link := dataList[1]
					requestId := dataList[2]
					method := dataList[3]
					reqTime := dataList[4]
					logger.Info("state = ", state, " | link = ", link, " | requestId = ", requestId, "| method", method, "| reqTime = ", reqTime)
				}

			}else{
				logger.Error("传入的数据太小或太大, 建议 1~10240个字节")
			}
		}

	}()
}