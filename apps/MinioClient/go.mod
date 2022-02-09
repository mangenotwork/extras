module github.com/mangenotwork/extras/apps/MinioClient

go 1.16

replace github.com/mangenotwork/extras/common => ../../common

require (
	github.com/mangenotwork/extras/common v0.0.0-00010101000000-000000000000
	github.com/minio/minio-go/v6 v6.0.57
	github.com/nfnt/resize v0.0.0-20180221191011-83c6a9932646
)
