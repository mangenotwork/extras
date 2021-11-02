package engine

import (
	"encoding/json"
	"github.com/mangenotwork/extras/apps/Push/model"
	//jsoniter "github.com/json-iterator/go"
	"github.com/mangenotwork/extras/apps/Push/mq"
	"log"
)

func StartMqServer(){
	log.Println("StartMqServer  启动消费者")

	go func() {
		var e = make(chan []byte)
		mq.NewMQ().Consumer("mange-push-send", e, sendMessage)
	}()

	go func() {
		var add = make(chan []byte)
		mq.NewMQ().Consumer("mange-push-device", add, deviceDo)
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
func deviceDo(b []byte) {
	log.Println("消费消息 : ", string(b))
	deviceData := &mq.MQDevice{}
	err := json.Unmarshal(b, &deviceData)
	if err != nil {
		log.Println("序列化错误")
		return
	}
	device := &model.Device{
		ID: deviceData.Device,
	}
	if deviceData.Type == "add" {
		// 当前服务存在连接则加入到topic
		if conn, ok := model.AllWsClient[deviceData.Device]; ok {
			_=device.SubTopic(conn, deviceData.Topic)
		}

		if conn, ok := model.AllTcpClient[deviceData.Device]; ok {
			_=device.SubTopic(conn, deviceData.Topic)
		}

		if conn, ok := model.AllUdpClient[deviceData.Device]; ok {
			_=device.SubTopic(conn, deviceData.Topic)
		}

	}

	if deviceData.Type == "del" {
		_=device.CancelTopic(deviceData.Topic)
	}

}