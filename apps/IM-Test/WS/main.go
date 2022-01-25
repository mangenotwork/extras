package main

import (
	"github.com/mangenotwork/extras/common/logger"
	gt "github.com/mangenotwork/gathertool"
	"log"
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
	log.Println(wsc, err)

	go func(){

		time.Sleep(2*time.Second)
		err = wsc.Send([]byte(`{
				"cmd":"Auth",
				"data":{
					"device":"456"
				}
			}
			`))
		time.Sleep(2*time.Second)
		log.Println(err)

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