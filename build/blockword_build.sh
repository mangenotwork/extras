#! /bin/bash

cd ../apps/BlockWord/

# app name
AppName=block-word

# app version
VERSION=0.0.3

# ImageURL
#ImageURL=registry.cn-shenzhen.aliyuncs.com/niupp/

# go mod
rm -rf vendor
go mod vendor


#builder app
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-w -s" -o $AppName main.go

# docker build
docker build --rm -t $AppName:latest .

# rm tmp file
if [ $? -eq 0 ];then
# rm tmp file
rm -rf AppName
rm -rf vendor
echo "publish:success"
else
echo "publish:failure"
fi