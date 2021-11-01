package engine

import (
	"encoding/json"
	"github.com/mangenotwork/extras/apps/Push/model"
	//jsoniter "github.com/json-iterator/go"
	"github.com/mangenotwork/extras/apps/Push/mq"
	"log"
)

func StartMqServer(){
	log.Println("StartMqServer")

	go func() {
		var e = make(chan []byte)
		mq.NewMQ().Consumer("mange-push-send", e, sendMessage)
	}()

	go func() {
		var add = make(chan []byte)
		mq.NewMQ().Consumer("mange-add-device", add, addDevice)
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
		topic.Send(sendData)
	}

}

// addDevice 添加设备
func addDevice(b []byte) {
	log.Println("消费消息 : ", string(b))
	deviceData := &mq.MQAddDevice{}
	err := json.Unmarshal(b, &deviceData)
	if err != nil {
		log.Println("序列化错误")
		return
	}
	device := &model.Device{
		ID: deviceData.Device,
	}
	// 当前服务存在连接则加入到topic
	conn, ok := model.AllWsClient[deviceData.Device]
	if ok {
		log.Println("存在连接 加入连接")
		_=device.SubTopic(conn, deviceData.Topic)
	}
}