## 推送服务
> 支持水平扩展, 支持http,udp,websocket协议的推送服务,主要模式有发布订阅,组播等; 提供 http/s,grpc api

## 配置说明
- app.runType 运行模式; dev:开发模式; prod:生产模式； test:测试模式；
- httpServer.open  是否启动http服务
- httpServer.prod  http服务端口
- grpcServer.open  是否启动grpc服务
- grpcServer.prod  grpc服务端口
- redis  redis配置
- tcpServer   tcp服务配置
- udpServer   udp服务配置
- mqType    消息队列类型, 目前支持 nsq, rabbit
- nsq   nsq客户端配置
- rabbit  rabbit客户端配置

```shell script
app:
  name: BlockWord
  runType: dev

httpServer:
  open: true
  prod: 1241

grpcServer:
  open: false
  prod: 1242
  log: true

tcpServer:
  open: true
  prod: 1243

udpServer:
  open: true
  prod: 1244

redis:
  host: 192.168.0.197
  port: 6379
  db: 4
  password: love0021$%
  maxidle: 50
  maxactive: 10000

mqType: rabbit

nsq:
  producer: 192.168.0.197:4150
  consumer: 192.168.0.197:4161

rabbit:
  addr: 192.168.0.191:5672/admin
  user: admin
  password: admin

```

## Http 文档

##### [post|get] /register
> 登记, 下发一个随机uuid可以作为设备id,以便确认设备

##### [post] /topic/create
> 创建 Topic

##### [post] /topic/publish
> 发布推送

##### [post] /topic/sub
> 设备订阅推送, 支持批量

##### [post] /topic/cancel
> 设备取消订阅, 支持批量



## WebSocket 文档


## TCP 文档


## UDP 文档


## grpc 文档
> proto文件: 

> 生成pb.go: 


## 编译
> 直接编译:  go build main.go

> 编译为docker: https://github.com/mangenotwork/extras/build/push_build.sh