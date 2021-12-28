module github.com/mangenotwork/extras/apps/Push

go 1.15

replace github.com/mangenotwork/extras/common => ../../common

require (
	github.com/google/uuid v1.1.2
	github.com/gorilla/websocket v1.4.2
	github.com/mangenotwork/extras/common v0.0.0-00010101000000-000000000000
	github.com/nsqio/go-nsq v1.1.0
	github.com/streadway/amqp v1.0.0
	go.mongodb.org/mongo-driver v1.7.3
	golang.org/x/net v0.0.0-20211019232329-c6ed85c7a12d
)
