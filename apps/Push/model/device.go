package model

import (
	"errors"
	"fmt"
	"github.com/mangenotwork/extras/common/rediscmd"
	"github.com/mangenotwork/extras/common/utils"
	"log"
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
func (d *Device) GetTopic(conn Client) {
	for _,v := range rediscmd.SMEMBERSString(fmt.Sprintf(DeviceTopic, d.ID)) {
		log.Println("获取订阅, 并加入订阅", v)
		topic, ok := TopicMap[v]
		if !ok {
			TopicMap[v] = &Topic{
				Name : v,
				ID : v,
				WsClient : make(map[string]*WsClient),
				TcpClient : make(map[string]*TcpClient),
				UdpClient : make(map[string]*UdpClient),
			}
			topic = TopicMap[v]
		}

		if conn.GetWsConn() != nil {
			topic.WsClient[d.ID] = conn.GetWsConn()
		}

		if conn.GetTcpConn() != nil {
			topic.TcpClient[d.ID] = conn.GetTcpConn()
		}

		if conn.GetUdpConn() != nil {
			topic.UdpClient[d.ID] = conn.GetUdpConn()
		}
	}
}

// 释放连接
func (d *Device) Discharge(connType string) {
	for _,v := range rediscmd.SMEMBERS(fmt.Sprintf(DeviceTopic, d.ID)) {
		log.Println(v)
		if topic, ok := TopicMap[utils.Any2String(v)]; ok {
			if connType == "ws" {
				delete(topic.WsClient, d.ID)
			}
			if connType == "tcp" {
				delete(topic.TcpClient, d.ID)
			}
			if connType == "udp" {
				delete(topic.UdpClient, d.ID)
			}
		}
	}

	for _,v := range rediscmd.SMEMBERS(fmt.Sprintf(DeviceGroup, d.ID)) {
		log.Println(v)
		if group, ok := GroupMap[v.(string)]; ok {
			if connType == "ws" {
				delete(group.WsClient, d.ID)
			}
			if connType == "tcp" {
				delete(group.TcpClient, d.ID)
			}
			if connType == "udp" {
				delete(group.UdpClient, d.ID)
			}
		}
	}
}

// 获取组
func (d *Device) GetGroup(conn Client) {
	for _,v := range rediscmd.SMEMBERS(fmt.Sprintf(DeviceGroup, d.ID)) {
		log.Println(v)
		group, ok := GroupMap[v.(string)]
		if !ok {
			GroupMap[v.(string)] = &Group{
				Name : v.(string),
				ID : v.(string),
				WsClient : make(map[string]*WsClient),
				TcpClient : make(map[string]*TcpClient),
				UdpClient : make(map[string]*UdpClient),
			}
			group = GroupMap[v.(string)]
		}

		if conn.GetWsConn() != nil {
			group.WsClient[d.ID] = conn.GetWsConn()
		}

		if conn.GetTcpConn() != nil {
			group.TcpClient[d.ID] = conn.GetTcpConn()
		}

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

func (d *Device) SubTopic(conn Client, topicName string) (err error) {
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
				TcpClient : make(map[string]*TcpClient),
				UdpClient : make(map[string]*UdpClient),
			}
			topic = TopicMap[topicName]
		}

		if conn.GetWsConn() != nil {
			topic.WsClient[d.ID] = conn.GetWsConn()
		}

		if conn.GetTcpConn() != nil {
			topic.TcpClient[d.ID] = conn.GetTcpConn()
		}

		if conn.GetUdpConn() != nil {
			topic.UdpClient[d.ID] = conn.GetUdpConn()
		}
	}
	return
}

func (d *Device) CancelTopic(topicName string) (err error) {
	if !rediscmd.EXISTS(fmt.Sprintf(TopicKey, topicName)) {
		err = errors.New("订阅的topic不存在!")
		return
	}
	if topic, ok := TopicMap[topicName]; ok {
		if _,ok := topic.WsClient[d.ID]; ok {
			delete(topic.WsClient, d.ID)
		}
		if _,ok := topic.TcpClient[d.ID]; ok {
			delete(topic.TcpClient, d.ID)
		}
		if _,ok := topic.UdpClient[d.ID]; ok {
			delete(topic.UdpClient, d.ID)
		}
	}
	err = rediscmd.SREMOne(fmt.Sprintf(DeviceTopic, d.ID), topicName)
	if err == nil {
		log.Println(err)
	}
	return
}
