package engine

import (
	"github.com/mangenotwork/extras/common/conf"
	"github.com/mangenotwork/extras/common/utils"
	"log"
	"net"
)

func StartUdpServer(){
	go func() {
		log.Println("StartUdpServer")

		// 监听
		listener, err := net.ListenUDP("udp", &net.UDPAddr{
			IP: net.ParseIP("0.0.0.0"),
			Port: utils.Str2Int(conf.Arg.UdpServer.Prod),
		})
		if err != nil {
			log.Println(err)
			return
		}
		log.Printf("Local: <%s> \n", listener.LocalAddr().String())

		// 读取数据
		data := make([]byte, 1024)
		for {
			n, remoteAddr, err := listener.ReadFromUDP(data)
			if err != nil {
				log.Printf("error during read: %s", err)
			}
			log.Printf("<%s> %s\n", remoteAddr, data[:n])
			_, err = listener.WriteToUDP([]byte("world"), remoteAddr)
			if err != nil {
				log.Printf(err.Error())
			}
		}

	}()
}