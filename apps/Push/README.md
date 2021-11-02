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

返回
```
{"code":0,"timestamp":1635842677,"msg":"succeed","data":"76b50d7f-9e0f-4d4f-915d-7782b325f909"}
```
---

##### [post] /topic/create
> 创建 Topic

参数
- topic_name 要创建的topic的name
```
{
    "name":"test"
}
```

返回
```
{
    "code": 0,
    "timestamp": 1635842830,
    "msg": "succeed",
    "data": "创建成功"
}
```
---

##### [post] /topic/publish
> 发布推送

参数
- topic_name  topic的name
- data  推送的数据
```
{
	"name": "t3",
	"data":"903ac7d9da690f831f10a78f6eff87ae"
}
```

返回
```
{
    "code": 0,
    "timestamp": 1635841971,
    "msg": "succeed",
    "data": "发送成功"
}
```
---

##### [post] /topic/sub
> 设备订阅推送, 支持批量

参数
- topic_name  topic的name
- device_list device列表
```
{
	"topic_name":"t3",
	"device_list": ["123", "456"]
}
```

返回
```
{
    "code": 0,
    "timestamp": 1635842123,
    "msg": "succeed",
    "data": "订阅成功"
}
```
---

##### [post] /topic/cancel
> 设备取消订阅, 支持批量

参数
- topic_name  topic的name
- device_list device列表
```
{
	"topic_name":"t3",
	"device_list": ["123", "456"]
}
```

返回
```
{
    "code": 0,
    "timestamp": 1635842123,
    "msg": "succeed",
    "data": "取消订阅成功"
}
```
---

##### [get] /device/view/topic 
> 查询设备订阅的topic

参数
- device  设备id
```
/device/view/topic?device=123
```

返回
```
{"code":0,"timestamp":1635842500,"msg":"succeed","data":["t1","t2","t3","test1"]}
```

---

##### [get] /topic/all/device
> 查询topic被哪些设备订阅

参数
- topic  TopicName
```
/topic/all/device?topic=t3
```

返回
```
{"code":0,"timestamp":1635842573,"msg":"succeed","data":["123","456"]}
```
--- 

##### [get] /topic/check/device
> 查询topic是否被指定device订阅

参数
- device  设备id
- topic  TopicName
```
/topic/check/device?device=123&topic=t3
```

返回
```
{"code":0,"timestamp":1635842628,"msg":"succeed","data":"123订阅了t3"}
```
---

##### [get] /topic/disconnection/all
> 强制指定topic下全部设备断开接收推送

参数
- topic  TopicName
```
/topic/disconnection/all?topic=t3
```

返回
```
{"code":0,"timestamp":1635844695,"msg":"succeed","data":"断开成功"}
```
---

## WebSocket 文档


## TCP 文档


## UDP 文档


## grpc 文档
> proto文件: 

> 生成pb.go: 


## 编译
> 直接编译:  go build main.go

> 编译为docker: https://github.com/mangenotwork/extras/build/push_build.sh