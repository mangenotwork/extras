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

	go func(){

		time.Sleep(1*time.Second)
		err = wsc.Send([]byte(`
{
	"cmd":"Auth",
	"data":{
		"device":"456"
	}
}
`))
		time.Sleep(1*time.Second)
		log.Println(err)


//		err = wsc.Send([]byte(`
//{
//	"cmd":"TopicJoin",
//	"data":{
//		"topic":"test1"
//	}
//}
//`))
//		log.Println(err)


	}()

	for {
		data := make([]byte,1024)
		err = wsc.Read(data)
		log.Println(err)
		log.Println("data = ", string(data))
		err = wsc.Send([]byte(``))
		log.Println(err)
	}

	wsc.Close()
}
