#! /bin/bash
cd ../api/WordHelper_Proto/
protoc --go_out=plugins=grpc:. *.proto
cp -rf proto/wordhelper.pb.go ../../apps/WordHelper/proto/wordhelper.pb.go
mkdir ../../apps/GrpcClient/wordhelper
mkdir ../../apps/GrpcClient/wordhelper/proto/
cp -rf proto/wordhelper.pb.go ../../apps/GrpcClient/wordhelper/proto/wordhelper.pb.go
rm proto/wordhelper.pb.go
rmdir proto
echo "生成 wordhelper.pb.go 完成"