#! /bin/bash
cd ../api/ImgHelper_Proto/
protoc --go_out=plugins=grpc:. *.proto
cp -rf proto/imghelper.pb.go ../../apps/WordHelper/proto/imghelper.pb.go
mkdir ../../apps/GrpcClient/imghelper
mkdir ../../apps/GrpcClient/imghelper/proto/
cp -rf proto/imghelper.pb.go ../../apps/GrpcClient/imghelper/proto/imghelper.pb.go
rm proto/imghelper.pb.go
rmdir proto
echo "生成 imghelper.pb.go 完成"