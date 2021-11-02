package service

import (
	"github.com/mangenotwork/extras/apps/Push/model"
	"log"
)

func Into(client model.Client, deviceId string) (device *model.Device) {
	device = &model.Device{
		ID: deviceId,
	}

	if device.OnLineState() {
		client.SendMessage("设备已经在线")
		return
	}
	client.IntoAllClient(deviceId)

	device.UpLine()
	// 获取 device 订阅过的所有 topic 并加入
	device.GetTopic(client)
	// 获取 device 加入的所有组 group
	device.GetGroup(client)
	client.SendMessage("连接成功")
	return
}

func Interactive(data *model.CmdData, client model.Client) (device *model.Device){
	switch data.Cmd {
	case "Auth":
		log.Println("Auth")
		// 设备认证
		if deviceData,ok := data.Data.(map[string]interface{})["device"]; ok {
			if deviceId, yes := deviceData.(string); yes && len(deviceId)>0 {
				log.Println("device id = ", deviceId)
				device = Into(client, deviceId)
			}
		}

	case "TopicJoin":
		// 订阅 TopicJoin
		if topicData,ok := data.Data.(map[string]interface{})["topic"]; ok {
			if topic, yes := topicData.(string); yes && len(topic)>0 {
				log.Println("topic = ", topic)
				err := device.SubTopic(client, topic)
				if err != nil {
					client.SendMessage(err.Error())
				}else{
					client.SendMessage("订阅成功")
				}
			}
		}

	case "TopicCancel":
		// 取消订阅 TopicCancel
		if topicData,ok := data.Data.(map[string]interface{})["topic"]; ok {
			if topic, yes := topicData.(string); yes && len(topic)>0 {
				log.Println("topic = ", topic)
				err := device.CancelTopic(topic)
				if err != nil {
					client.SendMessage(err.Error())
				}else{
					client.SendMessage("取消订阅成功")
				}
			}
		}

	case "GroupJoin":
		// 加入组

	case "GroupQuit":
		//退出组

	default:
		client.SendMessage("未知Cmd")

	}
	return device
}