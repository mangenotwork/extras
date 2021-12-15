#! /bin/bash
cd ../api/BlockWord_Proto/
protoc --plugin=protoc-gen-go=/media/mange/c73f23f4-81bc-4d93-961c-bb6643e59ea6/MyGo/bin/protoc-gen-go --go_out=plugins=grpc:. *.proto
cp -rf proto/blockword.pb.go ../../apps/BlockWord/proto/blockword.pb.go
cp -rf proto/blockword.pb.go ../../apps/GrpcClient/proto/blockword.pb.go
rm proto/blockword.pb.go
rmdir proto
echo "生成 blockword.pb.go 完成"