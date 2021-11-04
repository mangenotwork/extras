module github.com/mangenotwork/extras/apps/BlockWord

go 1.15

replace github.com/mangenotwork/extras/common => ../../common

require (
	github.com/garyburd/redigo v1.6.2
	github.com/gomodule/redigo v1.8.5 // indirect
	github.com/json-iterator/go v1.1.12
	github.com/mangenotwork/extras/common v0.0.0-00010101000000-000000000000
	go.mongodb.org/mongo-driver v1.7.3 // indirect
	golang.org/x/net v0.0.0-20211019232329-c6ed85c7a12d
	golang.org/x/time v0.0.0-20210723032227-1f47c861a9ac // indirect
	google.golang.org/grpc v1.41.0
	google.golang.org/protobuf v1.27.1
)
