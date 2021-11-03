package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/mangenotwork/extras/apps/Push/model"
	"github.com/mangenotwork/extras/apps/Push/mq"
	"github.com/mangenotwork/extras/common/conn"
	"github.com/mangenotwork/extras/common/rediscmd"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/bson"
	"log"
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

type TopicMsgData struct {
	Message string `bson:"message"`
	SendTime string `bson:"send_time"`
}


func TopicSend(topicName, msg string) (err error) {
	if !TopicIsHave(topicName) {
		err = errors.New("topic不存在!")
		return
	}
	// 生产消息
	mqMsg := mq.MQMsg{
		Topic: topicName,
		Message: msg,
		SendTime: time.Now().Format("2006-01-02 15:04:05"),
	}
	b, _ := json.Marshal(&mqMsg)
	mq.NewMQ().Producer("mange-push-send", b)
	// 记录推送数据
	_, err = conn.GetMongoCollection("mange_push", topicName).InsertOne(context.TODO(), TopicMsgData{msg, mqMsg.SendTime});
	if err != nil {
		log.Println("记录推送数据 错误: ", err)
		return
	}
	return
}

// 返回的 获取推送数据
type GetTopicSendBody struct {
	List []*TopicMsgData `json:"list"`
	Count int64	`json:"count"`
	Page int64	`json:"page"`
	Limit int64	`json:"limit"`
}

// 获取推送数据
func GetTopicSend(topicName string, page, limit int64) (data *GetTopicSendBody, err error){
	data = &GetTopicSendBody{
		List: make([]*TopicMsgData, 0),
		Count: 0,
	}
	log.Println("topicName = ", topicName)
	collection := conn.GetMongoCollection("mange_push", topicName)
	data.Count, err = collection.CountDocuments(context.Background(), bson.D{})
	if err != nil {
		log.Println("collection.CountDocuments err = ", err)
		return
	}
	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}
	data.Page = page
	data.Limit = limit
	findOptions := options.Find()
	findOptions.SetSkip((page-1)*limit).SetLimit(limit) // 分页
	findOptions.SetSort(bson.D{{"send_time", -1}}) // 排序  1 正序;  -1 逆序
	cur, err := collection.Find(context.TODO(), bson.D{}, findOptions)
	if err != nil {
		log.Println("collection.Find err = ", err)
		return
	}
	for cur.Next(context.TODO()) {
		// create a value into which the single document can be decoded
		var elem TopicMsgData
		err := cur.Decode(&elem)
		if err != nil {
			log.Println(err)
		}
		data.List = append(data.List, &elem)
	}
	if err := cur.Err(); err != nil {
		log.Println(err)
	}
	// Close the cursor once finished
	_=cur.Close(context.TODO())
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

func TopicDisconnectionDevice(topicName string) (err error) {
	mqDevice := mq.MQDevice{
		Type: "disconnection",
		Topic: topicName,
	}
	b, _ := json.Marshal(&mqDevice)
	mq.NewMQ().Producer("mange-push-device", b)
	return
}