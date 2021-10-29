package engine

import (
	"encoding/json"
	"github.com/mangenotwork/extras/apps/Push/model"

	//jsoniter "github.com/json-iterator/go"
	"github.com/mangenotwork/extras/apps/Push/mq"
	"log"
)

func StartMqServer(){
	go func() {
		log.Println("StartMqServer")

		var (
			e = make(chan []byte)
			//json = jsoniter.ConfigCompatibleWithStandardLibrary
		)

		mq.NewMQ().Consumer("mange-push-send", e, sendMessage)


	}()
}

// sendMessage 实现消息发送业务
func sendMessage(b []byte) {

	log.Println("消费消息 : ", string(b))


	sendData := &mq.MQMsg{}
	err := json.Unmarshal(b, &sendData)
	if err != nil {
		log.Println("序列化错误")
		return
	}

	if topic, ok := model.TopicMap[sendData.Topic]; ok {
		topic.Send(sendData.Data)
	}

}
