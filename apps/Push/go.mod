module github.com/mangenotwork/extras/apps/Push

go 1.15

replace github.com/mangenotwork/extras/common => ../../common

require (
	github.com/garyburd/redigo v1.6.2 // indirect
	github.com/gomodule/redigo v1.8.5 // indirect
	github.com/google/uuid v1.1.2
	github.com/gorilla/websocket v1.4.2
	github.com/mangenotwork/extras/common v0.0.0-00010101000000-000000000000
	github.com/nsqio/go-nsq v1.1.0
	github.com/streadway/amqp v1.0.0
	golang.org/x/net v0.0.0-20200822124328-c89045814202
	golang.org/x/time v0.0.0-20210723032227-1f47c861a9ac // indirect
	google.golang.org/grpc v1.41.0 // indirect
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b // indirect
)
