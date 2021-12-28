module github.com/mangenotwork/extras/apps/GrpcClient

go 1.15

replace github.com/mangenotwork/extras/common => ../../common

require (
	github.com/golang/protobuf v1.5.2
	github.com/mangenotwork/extras/common v0.0.0-00010101000000-000000000000
	golang.org/x/net v0.0.0-20211019232329-c6ed85c7a12d
	google.golang.org/grpc v1.42.0
)
