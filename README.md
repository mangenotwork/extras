## extras 临时演员
> 附加功能微服务,开箱即用。目的是扩展整个项目的功能性，提供api给前后端使用；

## 通讯协议支持
- http/s: 使用jsonp输出主要客户端使用
- grpc: 服务端远程调用
- tcp: 服务端使用
- udp: 服务端使用

## 服务

#### BlockWord 屏蔽词服务
> 屏蔽词增删该查，词语白名单等; 提供 http/s, grpc api

#### ShortLink 短链接服务
>  短链接生成,管理; 提供 http/s,grpc api

#### StringHelper 文字处理服务
> 重名验证， 分词， 自动生成姓名， 网名， 查询邮编， 查询旅游景点， 城市id查询， 拼音， 诗人查询， 诗词查询， 成语查询;
> 提供 http/s, grpc api

#### IM 即时聊天功能服务
> 即时聊天功能; 提供 websocket, tcp, udp

#### ImgHelper 图片功能服务
> 生成二维码, 图片压缩， 图片水印， 生成文字图片， 图片固定剪切， 生成gif， 图片固定拼接， 图片基础信息获取，固定旋转， 格式转换;
> 提供 http/s, grpc api


## LICENSE : MIT License



