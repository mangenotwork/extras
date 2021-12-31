module github.com/mangenotwork/extras/apps/BlockWord

go 1.15

replace github.com/mangenotwork/extras/common => ../../common

require (
	github.com/garyburd/redigo v1.6.3
	github.com/golang/protobuf v1.5.2
	github.com/mangenotwork/extras/common v0.0.0-00010101000000-000000000000
	golang.org/x/net v0.0.0-20211019232329-c6ed85c7a12d
	golang.org/x/time v0.0.0-20211116232009-f0f3c7e86c11 // indirect
	google.golang.org/grpc v1.42.0
)
