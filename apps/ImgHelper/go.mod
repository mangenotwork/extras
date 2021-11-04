module github.com/mangenotwork/extras/apps/ImgHelper

go 1.15

replace github.com/mangenotwork/extras/common => ../../common

require (
	github.com/boombuler/barcode v1.0.1
	github.com/mangenotwork/extras/common v0.0.0-00010101000000-000000000000
	github.com/nfnt/resize v0.0.0-20180221191011-83c6a9932646
	github.com/skip2/go-qrcode v0.0.0-20200617195104-da1b6568686e
	github.com/tuotoo/qrcode v0.0.0-20190222102259-ac9c44189bf2
	golang.org/x/net v0.0.0-20211101193420-4a448f8816b3
	golang.org/x/time v0.0.0-20210723032227-1f47c861a9ac // indirect
	google.golang.org/grpc v1.42.0 // indirect
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b // indirect
)
