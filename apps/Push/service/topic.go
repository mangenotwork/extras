package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/mangenotwork/extras/apps/Push/model"
	"github.com/mangenotwork/extras/apps/Push/mq"
	"github.com/mangenotwork/extras/common/rediscmd"
	"time"
)

func NewTopic(topicName string) (err error) {
	if rediscmd.EXISTS(fmt.Sprintf(model.TopicKey, topicName)) {
		err = errors.New("topic已经存在!")
		return
	}
	arg := make([]interface{},0)
	arg = append(arg, map[string]interface{}{"name":topicName})
	arg = append(arg, map[string]interface{}{"Creation":time.Now().Unix()})
	err = rediscmd.HMSET(fmt.Sprintf(model.TopicKey, topicName), arg)
	if err != nil {
		return
	}

	topic := &model.Topic{
		Name: topicName,
		ID: topicName,
		WsClient: make(map[string]*model.WsClient),
		TcpClient: make(map[string]*model.TcpClient),
		UdpClient: make(map[string]*model.UdpClient),
	}
	model.TopicMap[topicName] = topic
	return
}

func TopicSend(topicName, msg string) (err error) {
	if !TopicIsHave(topicName) {
		err = errors.New("topic不存在!")
		return
	}

	mqMsg := mq.MQMsg{
		Topic: topicName,
		Message: msg,
		SendTime: time.Now().Format("2006-01-02 15:04:05"),
	}
	b, _ := json.Marshal(&mqMsg)
	mq.NewMQ().Producer("mange-push-send", b)

	return
}

func TopicIsHave(topicName string) bool {
	return rediscmd.EXISTS(fmt.Sprintf(model.TopicKey, topicName))
}

func TopicAddDevice(topicName, device string) (err error) {
	mqDevice := mq.MQDevice{
		Type: "add",
		Topic: topicName,
		Device: device,
	}
	b, _ := json.Marshal(&mqDevice)
	mq.NewMQ().Producer("mange-push-device", b)
	return
}

func TopicDelDevice(topicName, device string) (err error) {
	mqDevice := mq.MQDevice{
		Type: "del",
		Topic: topicName,
		Device: device,
	}
	b, _ := json.Marshal(&mqDevice)
	mq.NewMQ().Producer("mange-push-device", b)
	return
}