## BlockWord 屏蔽词服务
> 屏蔽词增删查，词语白名单增删查; 提供 http/s, grpc api

## 环境变量配置
- RUNMODE 运行模式；dev:开发模式读取dev.yaml； prod:生产模式读取prod.yaml； test:测试模式读取test.yaml；
```shell script
export RUNMODE=dev
```

## 配置说明

- app.runType 运行模式; dev:开发模式; prod:生产模式； test:测试模式；
- httpServer.open  是否启动http服务
- httpServer.prod  http服务端口
- grpcServer.open  是否启动grpc服务
- grpcServer.prod  grpc服务端口
- redis  redis配置

```shell script
app:
  name: BlockWord
  runType: dev  

httpServer:
  open: true
  prod: 1231

grpcServer:
  open: true
  prod: 1232

redis:
  host: 192.168.0.197
  port: 6379
  db: 4
  password: root123
  maxidle: 50
  maxactive: 10000
```

## Http接口文档

####  [POST] /v1/do  
> 屏蔽词过滤

参数
```json
{
  "str":"我在路口交通进行口交就在这个路口交接",
  "sub":"???"
}
```
返回
```json
{"code":0,"timestamp":1635232884,"msg":"succeed","data":{"str":"我在路口交通进行???就在这个路口交接","sub":"???"}}
```

##### [GET] /v1/add?word= 
> 添加屏蔽词

参数
- word 添加的需要屏蔽的词语，如口交

##### [GET] /v1/del?word=  
> 删除屏蔽词

参数
- word 删除的屏蔽词语，如口交

##### [GET] /v1/list 
> 查看所有屏蔽词

返回
```json
{"code":0,"timestamp":1635240175,"msg":"succeed","data":["口交","废物"]}
```

##### [GET] /v1/white/add?word=  
> 词语白名单添加

参数
- word 添加不需要被屏蔽的词语，如路口； 假设屏蔽词为口交，“路口交通”不会被屏蔽为“路××通”

#####  [GET] /v1/white/del?word=   
> 词语白名单删除

参数
- word 删除不需要被屏蔽的词语，如路口; 假设屏蔽词为口交,“路口交通”会被屏蔽“

##### [GET] /v1/white/list 
> 查看所有词语白名单

返回
```json
{"code":0,"timestamp":1635240175,"msg":"succeed","data":["路口"]}
```

##### [POST|GET] /v1/ishave
> 是否存在非法词，应用场景：判断命名非法

参数 
- post : {"str":"我在路口交通进行口交就在这个路口交接"}
- get : /v1/ishave?str=我在路口交通进行口交就在这个路口交接

返回
```json
{"code":0,"timestamp":1635304798,"msg":"succeed","data":true}
```

##### [POST|GET] /v1/ishave/list
> 是否存在非法词如果存在返回所有屏蔽词，应用场景：判断命名非法,显示非法词

参数 
- post : {"str":"我在口交通进爱操你妈圣诞节在欧帕斯卡分废速度发完全二维卡"}
- get : /v1/ishave/list?str=我在口交通进爱操你妈圣诞节在欧帕斯卡分废速度发完全二维卡

返回
```json
{"code":0,"timestamp":1635306381,"msg":"succeed","data":["口交","操你妈"]}
```

## grpc
> proto文件: https://github.com/mangenotwork/extras/api/BlockWord_Proto/blockword.proto

> 生成pb.go: https://github.com/mangenotwork/extras/script/blockword_pb.sh

##### rpc Do (DoReq) returns (DoResp);
> 屏蔽词过滤
##### rpc Add (AddReq) returns (AddResp);
> 添加屏蔽词
##### rpc Del(DelReq) returns (DelResp);
> 删除屏蔽词
##### rpc Get(GetReq) returns (GetResp);
> 查看所有屏蔽词
##### rpc WhiteWordAdd (WhiteWordAddReq) returns (WhiteWordAddResp);
> 词语白名单添加
##### rpc WhiteWordDel(WhiteWordDelReq) returns (WhiteWordDelResp);
> 词语白名单删除
##### rpc WhiteWordGet(WhiteWordGetReq) returns (WhiteWordGetResp);
> 查看所有词语白名单
##### rpc IsHave(IsHaveReq) returns (IsHaveResp);
> 是否存在非法词，应用场景：判断命名非法
##### rpc IsHaveList(IsHaveListReq) returns (IsHaveListResp);
> 是否存在非法词如果存在返回所有屏蔽词，应用场景：判断命名非法,显示非法词

## 编译
> 直接编译:  go build main.go

> 编译为docker: https://github.com/mangenotwork/extras/build/blockword_build.sh

