package main

import (
	"encoding/json"
	"github.com/mangenotwork/extras/common/logger"
	"github.com/mangenotwork/extras/common/utils"
	gt "github.com/mangenotwork/gathertool"
	"time"
)

/*
	测试 TCP 的 client

*/

func main(){
	client := gt.NewTcpClient()
	client.Run("192.168.0.9:29123", w, f)
}

func f(client *gt.TcpClient){
	go func() {
		// 发送登录请求
		_,err := client.Send([]byte(`{
			"cmd":"Login",
			"data":{
				"account":"a11",
				"password":"123456",
				"device":"1",
				"source":"windows"
			}
		}`))
		if err != nil {
			logger.Debug("err = ", err)
		}
	}()
}

func w(client *gt.TcpClient, data []byte) {
	// 解析data
	//logger.Debug(string(data))
	cmdData := &CmdData{}
	jsonErr := json.Unmarshal(data, &cmdData)
	if jsonErr != nil {
		logger.Debug(jsonErr)
		return
	}

	if cmdData.Cmd == "Token" {
		// 发起心跳
		go func() {
			for {
					_,err := client.Send([]byte(`{
						"cmd":"HeartBeat",
						"data":"`+ utils.Any2String(cmdData.Data)+`"}`))
					if err != nil {
						client.RConn <- struct{}{}
						return
					}
					time.Sleep(2 * time.Second)

			}
		}()

	}
}

type CmdData struct {
	Cmd string `json:"cmd"`
	Data interface{} `json:"data"`  // obj
	Msg string `json:"msg"`
	Code int `json:"code"`
	Token string `json:"token"`
}
