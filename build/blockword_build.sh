#! /bin/bash

cd ../apps/BlockWord/

# app name
AppName=blockword

# app version
VERSION=0.0.1

# ImageURL
ImageURL=ccr.ccs.tencentyun.com/mange/

# go mod
rm -rf vendor
go mod vendor


#builder app
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-w -s" -o $AppName main.go

# docker build
docker build --rm -t $AppName:latest .

# docker push
docker login ccr.ccs.tencentyun.com --username=100015308690 --password=Lm_123456
docker tag $AppName:latest $ImageURL$AppName:$VERSION
docker push $ImageURL$AppName:$VERSION


# rm tmp file
if [ $? -eq 0 ];then
  # rm tmp file
  docker rmi $AppName:latest
  rm -rf $AppName
  rm -rf vendor
  echo "publish:success"
else
  echo "publish:failure"
fi