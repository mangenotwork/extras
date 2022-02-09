module github.com/mangenotwork/extras/apps/ShortLink

go 1.15

replace github.com/mangenotwork/extras/common => ../../common

require (
	github.com/garyburd/redigo v1.6.3
	github.com/mangenotwork/extras/common v0.0.0-00010101000000-000000000000
	google.golang.org/grpc v1.42.0
	google.golang.org/protobuf v1.26.0
)
