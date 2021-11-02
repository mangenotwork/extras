## 短链接服务
> 短链接生成,管理,短链接的黑白名单,短链接的时效性; 提供 http/s,grpc api

## 配置说明
- app.runType 运行模式; dev:开发模式; prod:生产模式； test:测试模式；
- httpServer.open  是否启动http服务
- httpServer.prod  http服务端口
- grpcServer.open  是否启动grpc服务
- grpcServer.prod  grpc服务端口
- redis  redis配置

```shell script
app:
  name: ShortLink
  runType: dev

httpServer:
  open: true
  prod: 1233

grpcServer:
  open: true
  prod: 1234
  log: true

redis:
  host: 192.168.0.197
  port: 6379
  db: 4
  password: love0021$%
  maxidle: 50
  maxactive: 10000
```

## Http接口文档

####  [post] /v1/add  
> 创建短链接

参数

|  参数名   | 类型  | 说明 |
|  ----  | ----  | ---- |
| url  | sting | 目的地址 |
| aging  | int |  时效，单位秒, -1表示不过期 |
| deadline  | int | 截止日期， 单位时间戳, 只有当aging为0时才用 |
| is_privacy  | bool | 是否隐私 |
| password  | string | 只有当is_privacy=true使用 |
| open_block_list  | bool | 是否启用黑名单，启用后黑名单不能访问 |
| block_list  | []string | 访问黑名单， OpenBlockList=true使用 |
| open_white_list  | bool | 是否启用白名单，启用后只有白名单才能访问 |
| white_list  | []string | 访问白名单， OpenWhiteList=true使用 |

传参例子:
```json
{
  "url" : "https://www.baidu.com"
}
```
返回
```json
{
    "code": 0,
    "timestamp": 1635818022,
    "msg": "succeed",
    "data": {
        "url": "/mwxrj1Fng",
        "password": "",
        "expire": "1970-01-01 08:00:00"
    }
}
```
注意:请求则需要加上当前host, 如 http://127.0.0.1:8080/mwxrj1Fng

##### [TODO] [post] /v1/modify
> 修改短链接

##### [TODO] [post] /v1/get   
> 获取短链接信息

##### [TODO] [post] /v1/del   
> 删除短链接


## grpc 文档
> proto文件: https://github.com/mangenotwork/extras/api/ShortLink_Proto/shortlink.proto

> 生成pb.go: https://github.com/mangenotwork/extras/script/shortlink.sh


## 编译
> 直接编译:  go build main.go

> 编译为docker: https://github.com/mangenotwork/extras/build/shortlink_build.sh
