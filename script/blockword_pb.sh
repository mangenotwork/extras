#! /bin/bash
cd ../api/BlockWord_Proto/
protoc --go_out=plugins=grpc:. *.proto
cp -rf proto/blockword.pb.go ../../apps/BlockWord/proto/blockword.pb.go
mkdir ../../apps/GrpcClient/blockword
mkdir ../../apps/GrpcClient/blockword/proto/
cp -rf proto/blockword.pb.go ../../apps/GrpcClient/blockword/proto/blockword.pb.go
rm proto/blockword.pb.go
rmdir proto
echo "生成 blockword.pb.go 完成"