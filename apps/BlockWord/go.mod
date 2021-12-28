module github.com/mangenotwork/extras/apps/BlockWord

go 1.15

replace github.com/mangenotwork/extras/common => ../../common

require (
	github.com/garyburd/redigo v1.6.3
	github.com/golang/protobuf v1.5.2
	github.com/mangenotwork/extras/common v0.0.0-00010101000000-000000000000
	go.etcd.io/etcd/client/v3 v3.5.1 // indirect
	go.mongodb.org/mongo-driver v1.7.3 // indirect
	golang.org/x/net v0.0.0-20211019232329-c6ed85c7a12d
	golang.org/x/time v0.0.0-20210723032227-1f47c861a9ac // indirect
	google.golang.org/grpc v1.42.0
)
