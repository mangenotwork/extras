package main

import (
	"github.com/mangenotwork/extras/common/logger"
	gt "github.com/mangenotwork/gathertool"
)

/*
	测试 UDP 的 client

*/

func main() {
	client := gt.NewUdpClient()
	client.Run("192.168.0.9", 29124, r, w)
}

func r(client *gt.UdpClient, data []byte) {
	logger.Debug(string(data))
	// TODO 下发了token, 就需要存储下来

}

func w(client *gt.UdpClient) {
	logger.Debug("send login")
	go func() {
		_, err := client.Send([]byte(`{
			"cmd":"Login",
			"data":{
				"account":"a11",
				"password":"123456",
				"device":"1",
				"source":"windows"
			}
		}`))
		logger.Error(err)
	}()
}
