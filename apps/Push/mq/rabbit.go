package mq

import (
	"fmt"
	"github.com/mangenotwork/extras/common/conf"
	"log"

	"github.com/streadway/amqp"
)

type MQRabbitService struct {
}

func (m *MQRabbitService) newProducer() string {
	log.Println("amqp://"+conf.Arg.Rabbit.User+":"+conf.Arg.Rabbit.Password+"@"+conf.Arg.Rabbit.Addr)
	return "amqp://"+conf.Arg.Rabbit.User+":"+conf.Arg.Rabbit.Password+"@"+conf.Arg.Rabbit.Addr
}


func (m *MQRabbitService) Producer(topic string, data []byte) {
	mq, err := newRabbitMQPubSub(topic, m.newProducer())
	if err != nil {
		log.Println("[rabbit]无法连接到队列")
		return
	}

	log.Println(fmt.Sprintf("[生产消息] topic : %s -->  %s", topic, string(data)))
	err = mq.publishPub(data)
	if err != nil {
		log.Println("[生产消息] 失败 ： " + err.Error())
	}
}

func (m *MQRabbitService) Consumer(topic string, ch chan []byte, f func(b []byte)) {
	mh, err := newRabbitMQPubSub(topic, m.newProducer())
	if err != nil {
		log.Println("[rabbit]无法连接到队列")
		return
	}
	msg := mh.registryReceiveSub()

	go func(m <-chan amqp.Delivery){
		for {
			select {
			case s := <-m:
				f(s.Body)
			}
		}
	}(msg)

	log.Println("[Rabbit] %v started", topic)
}

type rabbitMQ struct {
	//连接
	conn *amqp.Connection
	//管道
	channel *amqp.Channel
	//队列名称
	QueueName string
	//交换机
	Exchange string
	//key Simple模式 几乎用不到
	Key string
	//连接信息
	MqUrl string
}

//创建RabbitMQ结构体实例
func newRabbitMQ(queueName, exchange, key, mqUrl string) (*rabbitMQ, error) {
	mq := &rabbitMQ{QueueName: queueName, Exchange: exchange, Key: key, MqUrl: mqUrl}
	var err error
	//创建rabbitMq连接
	mq.conn, err = amqp.Dial(mq.MqUrl)
	if err != nil {
		log.Println("创建连接错误！", err)
		return nil, err
	}
	mq.channel, err = mq.conn.Channel()
	if err != nil {
		log.Println("获取channel失败", err)
	}
	return mq, err
}

//断开channel和connection
func (r *rabbitMQ) Destroy() {
	_=r.channel.Close()
	_=r.conn.Close()
}

//简单模式step：1。创建简单模式下RabbitMQ实例
func newRabbitMQSimple(queueName, mqUrl string) (*rabbitMQ, error) {
	return newRabbitMQ(queueName, "", "", mqUrl)
}

//简单模式Step:2、简单模式下生产代码
func (r *rabbitMQ) publishSimple(message []byte) (err error) {
	//1、申请队列，如果队列存在就跳过，不存在创建
	//优点：保证队列存在，消息能发送到队列中
	_, err = r.channel.QueueDeclare(
		r.QueueName, //队列名称
		false, //是否持久化
		false, //是否为自动删除 当最后一个消费者断开连接之后，是否把消息从队列中删除
		false, //是否具有排他性 true表示自己可见 其他用户不能访问
		false, //是否阻塞 true表示要等待服务器的响应
		nil, //额外参数
	)
	if err != nil {
		log.Println("[Rabbit] failed to declare a queue", err)
	}

	//2.发送消息到队列中
	err = r.channel.Publish(
		r.Exchange, //默认的Exchange交换机是default,类型是direct直接类型
		r.QueueName, //要赋值的队列名称
		false, //如果为true，根据exchange类型和rout key规则，如果无法找到符合条件的队列那么会把发送的消息返回给发送者
		false, //如果为true,当exchange发送消息到队列后发现队列上没有绑定消费者，则会把消息还给发送者
		amqp.Publishing{
			ContentType: "text/plain", //类型
			Body: message, //消息
		})
	if err != nil {
		log.Println("[Rabbit] publish 消息失败", err)
	}
	return
}

//简单模式注册消费者
func (r *rabbitMQ) registryConsumeSimple() (msg <-chan amqp.Delivery) {
	//1、申请队列，如果队列存在就跳过，不存在创建
	//优点：保证队列存在，消息能发送到队列中
	_, err := r.channel.QueueDeclare(
		//队列名称
		r.QueueName,
		//是否持久化
		false,
		//是否为自动删除 当最后一个消费者断开连接之后，是否把消息从队列中删除
		false,
		//是否具有排他性
		false,
		//是否阻塞
		false,
		//额外参数
		nil,
	)
	if err != nil {
		log.Println(err)
	}
	//接收消息
	msg, err = r.channel.Consume(
		r.QueueName,
		//用来区分多个消费者
		"",
		//是否自动应答
		true,
		//是否具有排他性
		false,
		//如果设置为true,表示不能同一个connection中发送的消息传递给这个connection中的消费者
		false,
		//队列是否阻塞
		false,
		nil,
	)
	if err != nil {
		log.Println(err)
	}
	return
}

//订阅模式创建 rabbitMq实例  (目前用的fanout模式)
func newRabbitMQPubSub(exchangeName, mqUrl string) (*rabbitMQ, error) {
	//创建rabbitMq实例
	mq, err := newRabbitMQ("", exchangeName, "", mqUrl)
	if mq == nil || err != nil {
		return nil, err
	}
	//获取connection
	mq.conn, err = amqp.Dial(mq.MqUrl)
	if mq.conn == nil || err != nil {
		return nil, err
	}
	//获取channel
	mq.channel, err = mq.conn.Channel()
	if err != nil {
		log.Println("[Rabbit] failed to open a channel!", err)
	}
	return mq, err
}

//订阅模式生成
func (r *rabbitMQ) publishPub(message []byte) (err error) {
	//尝试创建交换机，不存在创建
	err = r.channel.ExchangeDeclare(
		r.Exchange, //交换机名称
		amqp.ExchangeFanout, //交换机类型 广播类型
		true, //是否持久化
		false, //是否自动删除
		false, //true表示这个exchange不可以被client用来推送消息，仅用来进行exchange和exchange之间的绑定
		false, //是否阻塞 true表示要等待服务器的响应  false 无等待
		nil, //参数
	)
	if err != nil {
		log.Println("[Rabbit] failed to declare an exchange ",err)
	}

	//2 发送消息
	err = r.channel.Publish(
		r.Exchange,
		"",
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body: message,
		})
	return
}

//订阅模式消费端代码
func (r *rabbitMQ) registryReceiveSub() (msg <-chan amqp.Delivery) {
	//尝试创建交换机，不存在创建
	err := r.channel.ExchangeDeclare(
		r.Exchange, //交换机名称
		amqp.ExchangeFanout, //交换机类型 广播类型
		true, //是否持久化
		false, //是否字段删除
		false, //true表示这个exchange不可以被client用来推送消息，仅用来进行exchange和exchange之间的绑定
		false, //是否阻塞 true表示要等待服务器的响应
		nil,
	)
	if err != nil {
		log.Println("[Rabbit] failed to declare an exchange ",err)
	}

	//2试探性创建队列，创建队列
	q, err := r.channel.QueueDeclare(
		"", //随机生产队列名称
		false,
		false,
		true,
		false,
		nil,
	)
	if err != nil {
		log.Println("[Rabbit] Failed to declare a queue ", err)
	}

	//绑定队列到exchange中
	err = r.channel.QueueBind(
		q.Name,
		"", //在pub/sub模式下，这里的key要为空
		r.Exchange,
		false,
		nil,
	)
	//消费消息
	msg, err = r.channel.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	return
}

//话题模式 创建RabbitMQ实例
func newRabbitMQTopic(exchange, routingKey, mqUrl string) (*rabbitMQ, error) {
	mq, _ := newRabbitMQ("", exchange, routingKey, mqUrl)
	var err error
	mq.conn, err = amqp.Dial(mq.MqUrl)
	if err != nil {
		log.Println("[Rabbit] failed to connect rabbitMq! ", err)
	}
	mq.channel, err = mq.conn.Channel()
	if err != nil {
		log.Println("[Rabbit] failed to open a channel ", err)
	}
	return mq, err
}

//话题模式发送信息
func (r *rabbitMQ) publishTopic(message []byte) (err error) {
	err = r.channel.ExchangeDeclare(
		r.Exchange, //交换机名称
		amqp.ExchangeTopic, //交换机类型 话题模式
		true, //是否持久化
		false, //是否字段删除
		false, //true表示这个exchange不可以被client用来推送消息，仅用来进行exchange和exchange之间的绑定
		false, //是否阻塞 true表示要等待服务器的响应
		nil,
	)
	if err != nil {
		log.Println("[Rabbit] failed to declare an exchange ",err)
	}

	//2发送信息
	err = r.channel.Publish(
		r.Exchange,
		r.Key,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body: message,
		})
	return
}

//话题模式接收信息
//要注意key
//其中* 用于匹配一个单词，#用于匹配多个单词（可以是零个）
//匹配 xx.* 表示匹配xx.hello,但是xx.hello.one需要用xx.#才能匹配到
func (r *rabbitMQ) registryReceiveTopic() (msg <-chan amqp.Delivery) {
	//尝试创建交换机，不存在创建
	err := r.channel.ExchangeDeclare(
		r.Exchange, //交换机名称
		amqp.ExchangeTopic, //交换机类型 话题模式
		true, //是否持久化
		false, //是否字段删除
		false, //true表示这个exchange不可以被client用来推送消息，仅用来进行exchange和exchange之间的绑定
		false, //是否阻塞 true表示要等待服务器的响应
		nil,
	)
	if err != nil {
		log.Println("[Rabbit] failed to declare an exchange ",err)
	}

	//2试探性创建队列，创建队列
	q, err := r.channel.QueueDeclare(
		"", //随机生产队列名称
		false,
		false,
		true,
		false,
		nil,
	)
	if err != nil {
		log.Println("[Rabbit] Failed to declare a queue ", err)
	}

	//绑定队列到exchange中
	err = r.channel.QueueBind(
		q.Name,
		r.Key,
		r.Exchange,
		false,
		nil,
	)
	msg, err = r.channel.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	return
}

//路由模式 创建RabbitMQ实例
func newRabbitMQRouting(exchange, routingKey, mqUrl string) (*rabbitMQ, error) {
	rabbitMQ, _ := newRabbitMQ("", exchange, routingKey, mqUrl)
	var err error
	rabbitMQ.conn, err = amqp.Dial(rabbitMQ.MqUrl)
	if err != nil {
		log.Println("[Rabbit] failed   to connect rabbitMq! ", err)
	}
	rabbitMQ.channel, err = rabbitMQ.conn.Channel()
	if err != nil {
		log.Println("[Rabbit] failed to open a channel; ", err)
	}
	return rabbitMQ, err
}

//路由模式发送信息
func (r *rabbitMQ) publishRouting(message []byte) (err error) {
	err = r.channel.ExchangeDeclare(
		r.Exchange, //交换机名称
		amqp.ExchangeDirect, //交换机类型 广播类型
		true, //是否持久化
		false, 	//是否字段删除
		false, //true表示这个exchange不可以被client用来推送消息，仅用来进行exchange和exchange之间的绑定
		false, //是否阻塞 true表示要等待服务器的响应
		nil,
	)
	if err != nil {
		log.Println("[Rabbit] failed to declare an exchange ",err)
	}

	//发送信息
	err = r.channel.Publish(
		r.Exchange,
		r.Key,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body: message,
		})
	return
}

//路由模式接收信息
func (r *rabbitMQ) registryReceiveRouting() (msg <-chan amqp.Delivery) {
	err := r.channel.ExchangeDeclare(
		r.Exchange, //交换机名称
		amqp.ExchangeDirect, //交换机类型 广播类型
		true, //是否持久化
		false, //是否字段删除
		false, //true表示这个exchange不可以被client用来推送消息，仅用来进行exchange和exchange之间的绑定
		false, //是否阻塞 true表示要等待服务器的响应
		nil,
	)
	if err != nil {
		log.Println("[Rabbit] failed to declare an exchange ",err)
	}

	//2试探性创建队列，创建队列
	q, err := r.channel.QueueDeclare(
		"", //随机生产队列名称
		false,
		false,
		true,
		false,
		nil,
	)
	if err != nil {
		log.Println("[Rabbit] Failed to declare a queue ", err)
	}

	//绑定队列到exchange中
	err = r.channel.QueueBind(
		q.Name,
		r.Key,
		r.Exchange,
		false,
		nil,
	)
	msg, err = r.channel.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	return
}
