package main

import (
	"github.com/mangenotwork/extras/common/logger"
	gt "github.com/mangenotwork/gathertool"
	"log"
	"os"
	"time"
)

/*
	测试 WS 的 client

 */

func main(){
	// 请求  192.168.0.9:29127/login
	caseUrl := "http://192.168.0.9:29127/login"
	ctx, _ := gt.PostJson(caseUrl, `{"account":"a11","password":"123456"}`)
	rse := ctx.Json
	logger.Debug(rse)
	token := gt.Any2String(gt.Json2Map(rse)["data"])
	logger.Debug(token)

	// 连接ws
	host := "ws://192.168.0.9:29121"
	path := "/ws?token="+token+"&device=1&source=1"
	wsc, err := gt.WsClient(host, path, false)
	if err != nil {
		log.Println("连接失败 : ", err)
		os.Exit(0)
	}

	go func(){
		for {
			time.Sleep(2*time.Second)
			logger.Debug("send...")
			err = wsc.Send([]byte(`{
				"cmd":"HeartBeat",
				"data":"`+ token  +`"}`))
			log.Println(err)
		}
	}()

	for {
		data := make([]byte,1024)
		err = wsc.Read(data)
		log.Println(err)
		log.Println("data = ", string(data))
		err = wsc.Send([]byte(``))
		log.Println(err)
	}

}