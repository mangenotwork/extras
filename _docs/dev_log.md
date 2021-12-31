### 开发日志
- v0.0.18  [Push] 新增Push服务端engine实现
- v0.0.19  [新增PushClient] 新增PushClient
- v0.0.20  [Push] 新增ws连接断开客户端连接存储
- v0.0.21  [Push] 新增nsq消息队列
- v0.0.22  [Push] 新增rabbitMq消息队列
- v0.0.23  [Push] 新增订阅发布功能
- v0.0.24  [Push] 新增设备注册,设备订阅,设备取消订阅接口
- v0.0.25  [Push] 新增文档
- v0.0.26  [ShortLink] 新增文档

> v0.1
- v0.1.1  [push] 新增tcp连接断开客户端连接存储
- v0.1.2  [Push] 新增udp连接断开客户端连接存储
- v0.1.3  [Push] 新增查询设备订阅的topic, 查询topic被哪些设备订阅
- v0.1.4  [Push] 强制指定topic下全部设备断开接收推送
- v0.1.5  [Push] 引入 MongoDB
- v0.1.6  [Push] 记录推送数据
- v0.1.7  [ImgHelper] init project
- v0.1.8  [ImgHelper] 二维码生成, 条形码生成
- v0.1.9  [WordHelper] init project
- v0.1.10 [ConfigCenter] init project
- v0.1.11 [WordHelper] 分词接口
- v0.1.12 [WordHelper] Ocr接口
- v0.1.13 [WordHelper] 翻译接口, 文档更新
- v0.1.14 [WordHelper] pdf内容提取
- v0.1.15 [ImgHelper] 图片细节信息 exif
- v0.1.16 [WordHelper] 加密解密
- v0.1.17 [WordHelper] md转html
- v0.1.18 [ImgHelper] 图片压缩
- v0.1.19 [ImgHelper] 添加水印, 文字与图片水印
- v0.1.20 [ImgHelper] 生成文字图片, 应用场景: 文章转图片
- v0.1.21 [ImgHelper] 图片合成gif
- v0.1.22 [ImgHelper] 图片旋转
- v0.1.23 [ImgHelper] 图片居中(转为长宽一样的图片)
- v0.1.24 [ImgHelper] 图片拼接
- v0.1.25 [ImgHelper] 图片剪裁
- v0.1.26 [Push] 测试,改Bug,更新文档

> v0.2
- v0.2.1 [ImgHelper] 图片色彩反转
- v0.2.2 [ImgHelper] 图片灰化
- v0.2.3 [ImgHelper] 图片转为字符画
- v0.2.4 [ImgHelper] 图片透明
- v0.2.5 [ServiceTable] 由 ConfigCenter 改为 ServiceTable
- v0.2.6 [ServiceTable] 实现raft算法 - leader 选举
- v0.2.7 [ServiceTable] 配置文件优化
- v0.2.8 [ServiceTable] 设计数据存储方案
- v0.2.9 [ServiceTable] 数据结构集合
- v0.2.10 [ServiceTable] key增删查基于前缀树实现
- v0.2.11 [ServiceTable] 数据结构k/v
- v0.2.12 [ServiceTableClient] init project
- v0.2.13 [rpc] 新增rpc链路日志
- v0.2.14 [rpc] 新增rpc基于etcd的负载均衡
- v0.2.15 [WordHelper] 新增base64图片的识别
- v0.2.16 [WordHelper] OCR安装的文档
- v0.2.17 [IM-*] 文件结构初始化
- v0.2.18 [common] 增加日志
- v0.2.19 [common] 升级grpc和etcd版本
- v0.2.20 [LogCenter] 日志中心 文件结构初始化
- v0.2.21 [LogCenter] 收集日志
- v0.2.22 [common] 日志库上报日志
- v0.2.23 [common] 自定义http日志并上报
- v0.2.24 [common] 自定义grpc日志并上报
- v0.2.25 [LogCenter] 日志存储到 Boltdb
- v0.2.26 [common] http封装
- v0.3.1  [LogCenter] 获取grpc请求日志
- v0.3.2  [LogCenter] 获取http请求日志

> Todo
- v0.3.3 [LogCenter] 查看和下载日志文件
- v0.3.4
- v0.3.5
- v0.3.6
- v0.3.7
- v0.3.8
- v0.3.9
- v0.3.10

> 预计
- [ServiceTableClient] 请求设计
- [ServiceTableClient] 数据结构集合
- [ServiceTableClient] 租约
- [ServiceTableClient] 数据结构k/v
- [ServiceTable] 数据结构集合 - 分布式一致性
- [ServiceTable] 数据结构k/v - 分布式一致性
- [ShortLink] 新增查看短链接,修改短链接,删除短链接
- [Push] 指定设备强制断开
- [Push] 获取设备在线情况
- [Push] 设备拉取离线推送
- [Push] 设备接收推送的反馈
- [Push] 设备心跳,处理客户端心跳
- [Push] 设备接收推送的反馈
- [Push] 测试,改Bug,更新文档
- [WordHelper] 分词,Ocr,翻译 Grpc 
- [WordHelper] pdf内容提取,加密解密,md转html Grpc
- [ImgHelper] 生成数字或文字图的base64,每次的base64都不一样

> 研究
- 文本内容的领域信息识别, 基金数据, 两个文本之间的相似度, 开奖号抓取, 标签提取, 
- [WordHelper] html转md

