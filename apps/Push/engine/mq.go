package engine

import (
	"encoding/json"
	"github.com/mangenotwork/extras/apps/Push/model"
	//jsoniter "github.com/json-iterator/go"
	"github.com/mangenotwork/extras/apps/Push/mq"
	"github.com/mangenotwork/extras/common/logger"
)

func StartMqServer(){
	logger.Info("StartMqServer  启动消费者")

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
	logger.Info("消费消息 : ", string(b))
	sendData := &mq.MQMsg{}
	err := json.Unmarshal(b, &sendData)
	if err != nil {
		logger.Error("序列化错误")
		return
	}
	if topic, ok := model.TopicMap[sendData.Topic]; ok {
		topic.Send(sendData)
	}
}

// deviceDo 设备操作
func deviceDo(b []byte) {
	logger.Info("消费消息 : ", string(b))
	deviceData := &mq.MQDevice{}
	err := json.Unmarshal(b, &deviceData)
	if err != nil {
		logger.Error("序列化错误")
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

	if deviceData.Type == "disconnection" {
		model.TopicDisconnection(deviceData.Topic)
	}

}