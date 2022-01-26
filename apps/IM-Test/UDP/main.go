package main

import (
	"encoding/json"
	"github.com/mangenotwork/extras/common/logger"
	"github.com/mangenotwork/extras/common/utils"
	gt "github.com/mangenotwork/gathertool"
	"time"
)

/*
	测试 UDP 的 client

*/

var token = ""

func main() {
	client := gt.NewUdpClient()
	client.Run("192.168.0.9", 29124, r, w)
}

func r(client *gt.UdpClient, data []byte) {
	logger.Debug(string(data))

	cmdData := &CmdData{}
	jsonErr := json.Unmarshal(data, &cmdData)
	if jsonErr != nil {
		logger.Debug(jsonErr)
		return
	}

	// 下发了token, 就需要存储下来
	if cmdData.Cmd == "Token" {
		token = utils.Any2String(cmdData.Data)
		logger.Debug(token)
	}

	time.Sleep(2 * time.Second)
	_, err := client.Send([]byte(`{
			"cmd":"Hello",
			"token":"`+token+`"
		}`))
	logger.Debug(err)
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

type CmdData struct {
	Cmd string `json:"cmd"`
	Data interface{} `json:"data"`  // obj
	Msg string `json:"msg"`
	Code int `json:"code"`
	Token string `json:"token"`
}

