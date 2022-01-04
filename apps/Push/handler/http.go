package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/mangenotwork/extras/apps/Push/model"
	"github.com/mangenotwork/extras/apps/Push/service"
	"github.com/mangenotwork/extras/common/httpser"
	"github.com/mangenotwork/extras/common/logger"
	"github.com/mangenotwork/extras/common/middleware"
)

var upGrader = websocket.Upgrader{
	ReadBufferSize:  1024*100,
	WriteBufferSize: 65535,
	HandshakeTimeout: 5*time.Second,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func Hello(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_,_=w.Write([]byte("Hello ManGe"))
}

func Ws(w http.ResponseWriter, r *http.Request) {
	st := time.Now()
	var (
		device *model.Device
		client model.Client
		ip = middleware.GetIP(r)
	)

	deviceId := httpser.GetUrlArg(r, "device")
	conn, err := upGrader.Upgrade(w, r, nil)
	if err != nil {
		logger.Error("websocket upgrade error:%v", err)
		return
	}
	str := `
{
	"cmd":"Auth",
	"data":{
		"device":"*****"
	}
}
	`

	wsClient := model.NewWsClient().AddConn(conn).SetIP(ip)
	client = wsClient

	if len(deviceId) < 1 {
		_ = conn.WriteMessage(websocket.BinaryMessage, model.CmdDataMsg("[未知身份连接]device为空, 客户端可以send数据来确认身份: "+str))
	}else{
		device = service.Into(client, deviceId)
	}
	logger.Info("[连接日志] 连接成功. 用时 = ", time.Now().Sub(st))
	logger.Info("RemoteAddr = ", conn.RemoteAddr(), " | LocalAddr = ", conn.LocalAddr())

	for {
		_, data, err := conn.ReadMessage()
		if err != nil {
			// 释放客户端连接
			if device != nil {
				logger.Info("释放客户端连接")
				device.OffLine() // 下线记录
				delete(model.AllWsClient, deviceId)
				device.Discharge("ws") // 连接离开topic,group
			}
			_=conn.Close()
			return
		}
		logger.Info(data)
		if len(data) < 1 {
			continue
		}
		cmdData := &model.CmdData{}
		jsonErr := json.Unmarshal(data, &cmdData)
		if jsonErr != nil {
			_=conn.WriteMessage(websocket.BinaryMessage, model.CmdDataMsg("非法数据格式"))
			continue
		}
		logger.Info(cmdData)
		device = service.Interactive(cmdData, client)
	}

}

// 创建 Topic
type TopicCreateParam struct {
	Name string `json:"topic_name"`
}

func TopicCreate(w http.ResponseWriter, r *http.Request) {
	params := &TopicCreateParam{}
	httpser.GetJsonParam(r, params)
	err := service.NewTopic(params.Name)
	if err != nil {
		httpser.OutErrBody(w, 2001,err)
		return
	}
	httpser.OutSucceedBody(w, "创建成功")
}

// 发布
type PublishParam struct {
	TopicName string `json:"topic_name"`
	Data string  `json:"data"` // 发布的内容
}

func Publish(w http.ResponseWriter, r *http.Request) {
	params := &PublishParam{}
	httpser.GetJsonParam(r, params)
	err := service.TopicSend(params.TopicName, params.Data)
	if err != nil {
		httpser.OutErrBody(w, 2001,err)
		return
	}
	httpser.OutSucceedBody(w, "发送成功")
}

// 获取一个随机id, 可以作为设备id使用
func GetDeviceId(w http.ResponseWriter, r *http.Request) {
	httpser.OutSucceedBody(w, uuid.New().String())
}


type SubscriptionParam struct {
	TopicName string `json:"topic_name"`
	DeviceList []string  `json:"device_list"` // 发布的内容
}

// 设备订阅, 支持批量
func Subscription(w http.ResponseWriter, r *http.Request) {
	params := &SubscriptionParam{}
	httpser.GetJsonParam(r, params)

	if !service.TopicIsHave(params.TopicName) {
		httpser.OutErrBody(w, 2001, errors.New(params.TopicName + " Topic 不存在"))
		return
	}

	for _, v := range params.DeviceList {
		device := &model.Device{
			ID:v,
		}
		// 当前服务是否存在连接; 如果不存在则需要发布消息所有服务查找是否存在,如果存在则加入连接
		conn, ok := model.AllWsClient[v]
		if !ok {
			// 生产一条消息
			_=service.TopicAddDevice(params.TopicName, v)
		}
		_=device.SubTopic(conn, params.TopicName)
	}
	httpser.OutSucceedBody(w, "订阅成功")
}

// 设备取消订阅, 支持批量
func TopicCancel(w http.ResponseWriter, r *http.Request) {
	params := &SubscriptionParam{}
	httpser.GetJsonParam(r, params)
	if !service.TopicIsHave(params.TopicName) {
		httpser.OutErrBody(w, 2001, errors.New(params.TopicName + " Topic 不存在"))
		return
	}

	for _, v := range params.DeviceList {
		// 生产一条消息
		_=service.TopicDelDevice(params.TopicName, v)
	}
	httpser.OutSucceedBody(w, "取消订阅成功")
}

// 查询设备订阅的topic
func DeviceViewTopic(w http.ResponseWriter, r *http.Request) {
	device := httpser.GetUrlArg(r, "device")
	dataList := new(model.Device).SetId(device).AllTopic()
	httpser.OutSucceedBody(w, dataList)
}

// 查询topic被哪些设备订阅
func TopicAllDevice(w http.ResponseWriter, r *http.Request) {
	topic := httpser.GetUrlArg(r, "topic")
	dataList := model.GetTopicAllDevice(topic)
	httpser.OutSucceedBody(w, dataList)
}

// 查询topic是否被指定device订阅
func TopicCheckDevice(w http.ResponseWriter, r *http.Request) {
	device := httpser.GetUrlArg(r, "device")
	topic := httpser.GetUrlArg(r, "topic")
	isHas, _ := model.GetTopicHasDevice(topic, device)
	rse := device+"没有订阅"+topic
	if isHas {
		rse = device+"订阅了"+topic
	}
	httpser.OutSucceedBody(w, rse)
}

// 强制指定topic下全部设备断开接收推送
func TopicDisconnection(w http.ResponseWriter, r *http.Request) {
	topic := httpser.GetUrlArg(r, "topic")
	_= service.TopicDisconnectionDevice(topic)
	httpser.OutSucceedBody(w, "断开成功")
}

// 获取topic记录
func TopicLog(w http.ResponseWriter, r *http.Request) {
	topic := httpser.GetUrlArg(r, "topic")
	page := httpser.GetUrlArgInt64(r, "page")
	limit := httpser.GetUrlArgInt64(r, "limit")
	data, err := service.GetTopicSend(topic, page, limit)
	if err != nil {
		httpser.OutErrBody(w, 2001,err)
		return
	}
	httpser.OutSucceedBody(w, data)
}
