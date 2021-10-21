module github.com/mangenotwork/extras/apps/BlockWord

go 1.15

replace github.com/mangenotwork/extras/common => ../../common

require (
	github.com/mangenotwork/extras/common v0.0.0-00010101000000-000000000000
	golang.org/x/net v0.0.0-20211019232329-c6ed85c7a12d
	golang.org/x/time v0.0.0-20210723032227-1f47c861a9ac
	google.golang.org/grpc v1.41.0
	google.golang.org/protobuf v1.27.1
)
