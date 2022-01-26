package main

import (
	gt "github.com/mangenotwork/gathertool"
	"log"
)

/*
	测试 TCP 的 client

*/

func main(){
	client := gt.NewTcpClient()
	client.Run("192.168.0.9:29123", f)
}

func f(client *gt.TcpClient){
	go func() {
		// 发送登录请求
		_,err := client.Send([]byte(`{
			"cmd":"Login",
			"data":{
				"account":"a10",
				"password":"123456",
				"device":"1",
				"source":"windows"
			}
		}`))
		if err != nil {
			log.Println("err = ", err)
		}
	}()
}
