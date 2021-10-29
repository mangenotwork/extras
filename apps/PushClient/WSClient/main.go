package main

import (
	"log"
	"time"

	gt "github.com/mangenotwork/gathertool"
)

func main(){
	host := "ws://192.168.0.9:1241"
	path := "/ws"
	wsc, err := gt.WsClient(host, path)
	log.Println(wsc, err)
	for {
		err = wsc.Send([]byte(`okok`))
		log.Println(err)
		data := make([]byte,100)
		err = wsc.Read(data)
		log.Println(err)
		log.Println("data = ", string(data))
		time.Sleep(1*time.Second)
	}

	wsc.Close()
}
