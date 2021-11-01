package mq

import "github.com/mangenotwork/extras/common/conf"

type MQMsg struct {
	Topic string
	Data string
}

type MQ interface {
	Producer(topic string, data []byte)
	Consumer(topic string, ch chan []byte, f func(b []byte))
}

func NewMQ() MQ {
	switch conf.Arg.MqType {
	case "nsq":
		return new(MQNsqService)
	case "rabbit":
		return new(MQRabbitService)
	default:
		return new(MQNsqService)
	}
}