module github.com/mangenotwork/extras/apps/WordHelper

go 1.15

replace github.com/mangenotwork/extras/common => ../../common

require (
	github.com/mangenotwork/extras/common v0.0.0-00010101000000-000000000000
	github.com/otiai10/gosseract/v2 v2.3.1
	github.com/yanyiwu/gojieba v1.1.2
	golang.org/x/net v0.0.0-20211104170005-ce137452f963 // indirect
	golang.org/x/time v0.0.0-20210723032227-1f47c861a9ac // indirect
	google.golang.org/grpc v1.42.0 // indirect
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b // indirect
)
