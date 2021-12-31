module github.com/mangenotwork/extras/apps/WordHelper

go 1.15

replace github.com/mangenotwork/extras/common => ../../common

require (
	github.com/golang/protobuf v1.5.2
	github.com/mangenotwork/extras/common v0.0.0-00010101000000-000000000000
	github.com/microcosm-cc/bluemonday v1.0.16
	github.com/otiai10/gosseract/v2 v2.3.1
	github.com/russross/blackfriday v1.6.0
	github.com/yanyiwu/gojieba v1.1.2
	golang.org/x/crypto v0.0.0-20210513164829-c07d793c2f9a
	golang.org/x/net v0.0.0-20211104170005-ce137452f963
	google.golang.org/grpc v1.42.0
	google.golang.org/protobuf v1.26.0
)
