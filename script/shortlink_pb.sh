#! /bin/bash
cd ../api/ShortLink_Proto/
protoc --go_out=plugins=grpc:. *.proto
cp -rf proto/shortlink.pb.go ../../apps/ShortLink/proto/shortlink.pb.go
mkdir ../../apps/GrpcClient/shortlink
mkdir ../../apps/GrpcClient/shortlink/proto/
cp -rf proto/shortlink.pb.go ../../apps/GrpcClient/shortlink/proto/shortlink.pb.go
rm proto/shortlink.pb.go
rmdir proto
echo "生成 shortlink.pb.go 完成"