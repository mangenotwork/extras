package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/mangenotwork/extras/apps/Push/model"
	"github.com/mangenotwork/extras/apps/Push/mq"
	"github.com/mangenotwork/extras/common/rediscmd"
	"net"
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
		TcpClient: make(map[string]*net.Conn),
		UdpClient: make(map[string]*net.UDPAddr),
	}
	model.TopicMap[topicName] = topic
	return
}

func TopicSend(topicName, msg string) (err error) {
	if !rediscmd.EXISTS(fmt.Sprintf(model.TopicKey, topicName)) {
		err = errors.New("topic不存在!")
		return
	}

	mqMsg := mq.MQMsg{
		Topic: topicName,
		Data: msg,
	}
	b, _ := json.Marshal(&mqMsg)
	mq.NewMQ().Producer("mange-push-send", b)

	//if topic, ok := model.TopicMap[topicName]; ok {
	//	for _, v := range topic.WsClient {
	//		v.TopicSend(msg)
	//	}
	//}

	return
}