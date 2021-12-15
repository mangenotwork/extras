module github.com/mangenotwork/extras/apps/GrpcClient

go 1.15

replace github.com/mangenotwork/extras/common => ../../common

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0

require (
	github.com/golang/protobuf v1.5.0
	github.com/mangenotwork/extras/common v0.0.0-00010101000000-000000000000 // indirect
	github.com/zituocn/gow v1.0.9 // indirect
	golang.org/x/net v0.0.0-20211019232329-c6ed85c7a12d
	google.golang.org/grpc v1.41.0
)
