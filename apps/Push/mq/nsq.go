package mq

import (
	"fmt"

	"github.com/mangenotwork/extras/common/conf"
	"github.com/mangenotwork/extras/common/logger"
	"github.com/mangenotwork/extras/common/utils"
	gnsq "github.com/nsqio/go-nsq"
)

type MQNsqService struct {
}

func (m *MQNsqService) newProducer(addr string) (*gnsq.Producer, error) {
	if addr == "" {
		return nil, fmt.Errorf("[NSQ] init failed：need nsq server addr")
	}
	config := gnsq.NewConfig()
	p, err := gnsq.NewProducer(addr, config)
	p.SetLogger(nil, 1)
	return p, err
}

func (m *MQNsqService) Producer(topic string, data []byte) {
	client, err := m.newProducer(conf.Arg.Nsq.Producer)
	if err != nil {
		logger.Error("[nsq]无法连接到队列")
		return
	}
	logger.Info(fmt.Sprintf("[生产消息] topic : %s -->  %s", topic, string(data)))
	err = client.Publish(topic, data)
	if err != nil {
		logger.Error("[生产消息] 失败 ： " + err.Error())
	}
}

func (m *MQNsqService) Consumer(topic string, ch chan []byte, f func(b []byte)) {

	id,_ := utils.ID64()
	mh, err := newMessageHandler(conf.Arg.Nsq.Consumer, utils.Int642Str(id))
	if err != nil {
		logger.Error(err)
		return
	}
	go func() {
		mh.setMaxInFlight(1000)
		mh.registry(topic, ch)
	}()

	go func() {
		for {
			select {
			case s := <-ch:
				f(s)
			}
		}
	}()

	logger.Info("[NSQ] ServerID:%v => %v started", "mange-push", topic)
}

type messageHandler struct {
	msgChan   chan *gnsq.Message
	stop      bool
	nsqServer string
	Channel   string
	maxInFlight int
}

// NewMessageHandler return new MessageHandler
func newMessageHandler(nsqServer string, channel string) (mh *messageHandler, err error) {
	if nsqServer == "" {
		err = fmt.Errorf("[NSQ] need nsq server")
		return
	}
	mh = &messageHandler{
		msgChan:   make(chan *gnsq.Message, 1024),
		stop:      false,
		nsqServer: nsqServer,
		Channel:   channel,
	}

	return
}

// SetMaxInFlight set nsq consumer MaxInFlight
func (m *messageHandler) setMaxInFlight(val int){
	m.maxInFlight = val
}

// Registry registry nsq topic
func (m *messageHandler) registry(topic string, ch chan []byte) {
	config := gnsq.NewConfig()
	if m.maxInFlight > 0 {
		config.MaxInFlight = m.maxInFlight
	}
	consumer, err := gnsq.NewConsumer(topic, m.Channel, config)
	if err != nil {
		panic(err)
	}
	consumer.SetLogger(nil, 0)
	consumer.AddHandler(gnsq.HandlerFunc(m.handlerMessage))
	err = consumer.ConnectToNSQLookupd(m.nsqServer)
	if err != nil {
		panic(err)
	}
	m.process(ch)

}

// process process
func (m *messageHandler) process(ch chan<- []byte) {
	m.stop = false
	for {
		select {
		case message := <-m.msgChan:
			ch <- message.Body
			if m.stop {
				close(m.msgChan)
				return
			}
		}
	}
}

// handlerMessage handlerMessage
func (m *messageHandler) handlerMessage(message *gnsq.Message) error {
	if !m.stop {
		m.msgChan <- message
	}
	return nil
}
