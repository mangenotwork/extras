## 文本语言相关处理服务
> 分词， OCR, 翻译, 加密解密, 文本内容的领域信息, 文本相似度, 彩票开奖, 拼音, 标签提取, PDF提取
>
> 提供 http/s, grpc api

## 配置说明
- app.runType 运行模式; dev:开发模式; prod:生产模式； test:测试模式；
- httpServer.open  是否启动http服务
- httpServer.prod  http服务端口
- grpcServer.open  是否启动grpc服务
- grpcServer.prod  grpc服务端口

```shell script
app:
  name: WordHelper
  runType: dev

httpServer:
  open: true
  prod: 1271

grpcServer:
  open: true
  prod: 1272
  log: true
```

## Http接口文档

#### [get] /fenci/jieba
> 结巴分词

参数
- str 文本
- type 分词类型,如下

|  type 值   | 类型  |
|  ----  | ----   |
| 1  | 全模式  |
| 2  | 精确模式 |
| 3  | 搜索引擎模式 |
| 4  | 词性标注 |
| 5  | Tokenize 搜索引擎模式 |
| 6  | Tokenize 默认模式 |
| 7  | Extract |

返回
```json
{
  "code":0,
  "timestamp":1636424481,
  "msg":"succeed",
  "data":[
    {"Word":"日本京都大学","Weight":13.2075304714},
    {"Word":"计算所","Weight":11.7034530746},
    {"Word":"小明","Weight":11.1280889297},
    {"Word":"深造","Weight":9.0884932966},
    {"Word":"硕士","Weight":8.87023973058}
  ]
}
```
---

##### [get] /ocr/version
> 查看tesseract版本

返回
```json
{"code":0,"timestamp":1636439749,"msg":"succeed","data":"4.0.0"}
```

---

##### [get] /ocr/languages
> 查看tesseract支持的词语库

返回
```json
{"code":0,"timestamp":1636428427,"msg":"succeed","data":["chi_sim","chi_tra","deu","eng","jpn","osd"]}
```

---

##### [post] /ocr
> tesseract进行识别
>
> 注意:宿主机需要安装tesseract与chi_sim.traineddata词语库

参数(from-data)
- file 被识别的图片
- lang 识别的词语库,默认是chi_sim

返回
```json
{
    "code": 0,
    "timestamp": 1636428522,
    "msg": "succeed",
    "data": "预 分 配 切 片 和 际 射 #\n\n尝 试 皑 终 预 免 分 配 切 片 和 映 射 。 如 果 您 知 遵 要 兆 骋 放 置 的 元 素 数 春 , 请 使 用 诚 知 识 ! 这 春 着 改 善 了 武 类 代 团 的 运 迟 。\n将 此 视 为 微 化 , 但 是 , 始 终 这 样 做 是 一 积 很 好 的 横 式 , 因 为 它 不 会 增 加 太 多 复 杂 伯 。 性 能 方 面 , 它 仅 与 具 有 大 数 组\n的 关 镰 代 团 路 径 根 关 。\n\n注 意 ; 这 是 团 为 , 在 非 常 简 单 的 视 图 中 ,Go 运 行 时 分 配 当 前 大 小 的 2 俘 。 因 此 , 如 果 您 期 望 有 数 百 万 个 元 袁 ,Go\n将 ppend 在 两 者 之 闰 进 行 多 次 分 配 , 而 不 是 在 您 迹 行 预 分 配 时 只 进 行 一 次 分 配"
}
```

#### [get] /fanyi
> 翻译

参数
- word  需要翻译的词语

返回
```json
{"type":"ZH_CN2EN","errorCode":0,"elapsedTime":0,"translateResult":[[{"src":"你好","tgt":"hello"}]]}
```

---

## grpc 文档
> proto文件: https://github.com/mangenotwork/extras/api/WordHelper_Proto/wordhelper.proto

> 生成pb.go: https://github.com/mangenotwork/extras/script/wordhelper_pb.sh


## 编译
> 直接编译:  go build main.go

> 编译为docker: https://github.com/mangenotwork/extras/build/shortlink_build.sh
