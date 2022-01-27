#! /bin/bash

# IM-User
cd ../api/IM_Proto/user/
protoc --go_out=plugins=grpc:. *.proto
cp -rf proto/im-user.pb.go ../../../apps/IM-User/proto/im-user.pb.go
cp -rf proto/im-user.pb.go ../../../apps/IM-Conn/proto/im-user.pb.go
cp -rf proto/im-user.pb.go ../../../apps/IM-Msg/proto/im-user.pb.go
rm proto/im-user.pb.go
rmdir proto
echo "生成 im-user.pb.go 完成"

# IM-Msg
cd ../api/IM_Proto/msg/
protoc --go_out=plugins=grpc:. *.proto
cp -rf proto/im-msg.pb.go ../../../apps/IM-User/proto/im-msg.pb.go
cp -rf proto/im-msg.pb.go ../../../apps/IM-Conn/proto/im-msg.pb.go
cp -rf proto/im-msg.pb.go ../../../apps/IM-Msg/proto/im-msg.pb.go
rm proto/im-msg.pb.go
rmdir proto
echo "生成 im-msg.pb.go 完成"

# IM-Conn

# IM-Client

