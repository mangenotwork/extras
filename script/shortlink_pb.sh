#! /bin/bash
cd ../api/ShortLink_Proto/
protoc --go_out=plugins=grpc:. *.proto
cp -rf proto/shortlink.pb.go ../../apps/ShortLink/proto/blockword.pb.go
cp -rf proto/shortlink.pb.go ../../apps/GrpcClient/proto/shortlink.pb.go
rm proto/shortlink.pb.go
rmdir proto
echo "生成 shortlink.pb.go 完成"