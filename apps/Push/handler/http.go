package handler

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"github.com/mangenotwork/extras/apps/Push/model"
	"github.com/mangenotwork/extras/common/middleware"
	"github.com/mangenotwork/extras/common/utils"
	"log"
	"net/http"
	"time"
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
		wsUser *model.WsClient
		ip = middleware.GetIP(r)
	)

	deviceId := utils.GetUrlArg(r, "device")
	conn, err := upGrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("websocket upgrade error:%v", err)
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
	wsUser = &model.WsClient{
		Conn : conn,
		IP : ip,
	}

	if len(deviceId) < 1 {
		msg := &model.CmdData{
			Cmd: "message",
			Data: "[未知身份连接]device为空, 客户端可以send数据来确认身份: "+str,
		}
		_ = conn.WriteMessage(websocket.BinaryMessage, msg.Byte())
	}else{
		device = into(wsUser, deviceId, ip)
	}

	log.Println("[连接日志] 连接成功. 用时 = ", time.Now().Sub(st))
	log.Println("RemoteAddr = ", conn.RemoteAddr(), " | LocalAddr = ", conn.LocalAddr())

	for {
		_, data, err := conn.ReadMessage()
		if err != nil {
			// 释放客户端连接
			log.Println("释放客户端连接")
			device.OffLine() // 下线记录
			delete(model.AllWsClient, deviceId)
			device.Discharge() // 连接离开topic,group
			return
		}
		log.Println(data)
		if len(data) < 1 {
			continue
		}

		cmdData := &model.CmdData{}
		msg := &model.CmdData{}
		jsonErr := json.Unmarshal(data, &cmdData)
		if jsonErr != nil {
			msg.Cmd = "Message"
			msg.Data = "非法数据格式"
			_=conn.WriteMessage(websocket.BinaryMessage, msg.Byte())
			continue
		}
		log.Println(cmdData)
		switch cmdData.Cmd {
		case "Auth":
			// 设备认证
			if deviceData,ok := cmdData.Data.(map[string]interface{})["device"]; ok {
				if deviceId, yes := deviceData.(string); yes && len(deviceId)>0 {
					log.Println("device id = ", deviceId)
					device = into(wsUser, deviceId, ip)
				}
			}
		case "TopicJoin":
			// 订阅 TopicJoin
			if topicData,ok := cmdData.Data.(map[string]interface{})["topic"]; ok {
				if topic, yes := topicData.(string); yes && len(topic)>0 {
					log.Println("topic = ", topic)
					err := device.SubTopic(wsUser, topic)
					if err != nil {
						wsUser.SendMessage(err.Error())
					}else{
						wsUser.SendMessage("订阅成功")
					}
				}
			}
		case "TopicCancel":
			// 取消订阅 TopicCancel
			if topicData,ok := cmdData.Data.(map[string]interface{})["topic"]; ok {
				if topic, yes := topicData.(string); yes && len(topic)>0 {
					log.Println("topic = ", topic)
					err := device.CancelTopic(topic)
					if err != nil {
						wsUser.SendMessage(err.Error())
					}else{
						wsUser.SendMessage("取消订阅成功")
					}
				}
			}
		case "GroupJoin":
			// 加入组
		case "GroupQuit":
			//退出组
		default:
			msg.Cmd = "Message"
			msg.Data = "未知Cmd"
			_=conn.WriteMessage(websocket.BinaryMessage, msg.Byte())
		}

	}

}

func into(wsUser *model.WsClient, deviceId, ip string) (device *model.Device) {
	device = &model.Device{
		ID: deviceId,
	}

	if device.OnLineState() {
		wsUser.SendMessage("设备已经在线")
		return
	}
	model.AllWsClient[deviceId] = wsUser
	device.UpLine()
	// 获取 device 订阅过的所有 topic 并加入
	device.GetTopic(wsUser)
	// 获取 device 加入的所有组 group
	device.GetGroup(wsUser)
	wsUser.SendMessage("连接成功")
	return
}


// 创建 Topic
type TopicCreateParam struct {
	Name string
}

func TopicCreate(w http.ResponseWriter, r *http.Request) {

}

// 发布
type PublishParam struct {
	TopicName string
	Data string  // 发布的内容
}

func Publish(w http.ResponseWriter, r *http.Request) {

}

// 设备注册
func Register(w http.ResponseWriter, r *http.Request) {

}

// 设备订阅, 支持批量
func Subscription(w http.ResponseWriter, r *http.Request) {

}

// 设备取消订阅, 支持批量
func TopicCancel(w http.ResponseWriter, r *http.Request) {

}

