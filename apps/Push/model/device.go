package model

import (
	"fmt"
	"github.com/mangenotwork/extras/common/rediscmd"
	"github.com/mangenotwork/extras/common/utils"
	"log"
	"net"
	"errors"
)

type Device struct {
	ID string
}

const (
	DeviceTopic = "d:%s:topic" // 设备订阅的topic
	DeviceGroup = "d:%s:group" // 设备加入的组
	DeviceOnLine = "d:%s:onlien" // 0:不在线; 1:在线;
	TopicKey = "topic:%s" //
)

// 获取订阅, 并加入订阅
func (d *Device) GetTopic(conn *WsClient) {
	for _,v := range rediscmd.SMEMBERSString(fmt.Sprintf(DeviceTopic, d.ID)) {
		log.Println("获取订阅, 并加入订阅", v)
		topic, ok := TopicMap[v]
		if !ok {
			TopicMap[v] = &Topic{
				Name : v,
				ID : v,
				WsClient : make(map[string]*WsClient),
				TcpClient : make(map[string]*net.Conn),
				UdpClient : make(map[string]*net.UDPAddr),
			}
			topic = TopicMap[v]
		}
		topic.WsClient[d.ID] = conn
	}
}

// 释放连接
func (d *Device) Discharge() {
	for _,v := range rediscmd.SMEMBERS(fmt.Sprintf(DeviceTopic, d.ID)) {
		log.Println(v)
		if topic, ok := TopicMap[utils.Any2String(v)]; ok {
			delete(topic.WsClient, d.ID)
		}
	}

	for _,v := range rediscmd.SMEMBERS(fmt.Sprintf(DeviceGroup, d.ID)) {
		log.Println(v)
		if group, ok := GroupMap[v.(string)]; ok {
			delete(group.WsClient, d.ID)
		}
	}
}

// 获取组
func (d *Device) GetGroup(conn *WsClient) {
	for _,v := range rediscmd.SMEMBERS(fmt.Sprintf(DeviceGroup, d.ID)) {
		log.Println(v)
		group, ok := GroupMap[v.(string)]
		if !ok {
			GroupMap[v.(string)] = &Group{
				Name : v.(string),
				ID : v.(string),
				WsClient : make(map[string]*WsClient),
				TcpClient : make(map[string]*net.Conn),
				UdpClient : make(map[string]*net.UDPAddr),
			}
			group = GroupMap[v.(string)]
		}
		group.WsClient[d.ID] = conn
	}
}

// 上线
func (d *Device) UpLine() {
	_=rediscmd.SETEX(fmt.Sprintf(DeviceOnLine, d.ID), 60, 1)
}

// 下线
func (d *Device) OffLine() {
	log.Println(d)
	_=rediscmd.SETEX(fmt.Sprintf(DeviceOnLine, d.ID), 60, 0)
}

func (d *Device) OnLineState() bool {
	if rediscmd.Get(fmt.Sprintf(DeviceOnLine,d.ID)) == "1" {
		return true
	}
	return false
}

func (d *Device) SubTopic(conn *WsClient, topicName string) (err error) {
	if !rediscmd.EXISTS(fmt.Sprintf(TopicKey, topicName)) {
		err = errors.New("订阅的topic不存在!")
		return
	}
	err = rediscmd.SADD(fmt.Sprintf(DeviceTopic, d.ID), []interface{}{topicName})
	if err == nil {
		topic, ok := TopicMap[topicName]
		if !ok {
			TopicMap[topicName] = &Topic{
				Name : topicName,
				ID : topicName,
				WsClient : make(map[string]*WsClient),
				TcpClient : make(map[string]*net.Conn),
				UdpClient : make(map[string]*net.UDPAddr),
			}
			topic = TopicMap[topicName]
		}
		topic.WsClient[d.ID] = conn
	}
	return
}

func (d *Device) CancelTopic(topicName string) (err error) {
	if !rediscmd.EXISTS(fmt.Sprintf(TopicKey, topicName)) {
		err = errors.New("订阅的topic不存在!")
		return
	}
	if topic, ok := TopicMap[topicName]; ok {
		delete(topic.WsClient, d.ID)
	}
	err = rediscmd.SREMOne(fmt.Sprintf(DeviceTopic, d.ID), topicName)
	if err == nil {
		log.Println(err)
	}
	return
}