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

### tesseract-ocr 安装
- tesserocr GitHub：https://github.com/sirfz/tesserocr
- tesserocr PyPI：https://pypi.python.org/pypi/tesserocr
- tesseract下载地址：http://digi.bib.uni-mannheim.de/tesseract
- tesseract GitHub：https://github.com/tesseract-ocr/tesseract
- tesseract语言包：https://github.com/tesseract-ocr/tessdata
- tesseract文档：https://github.com/tesseract-ocr/tesseract/wiki/Documentation
- 下载词典: https://tesseract-ocr.github.io/tessdoc/Data-Files  
```
# Ubuntu、Debian和Deepin
$ sudo apt-get install -y tesseract-ocr libtesseract-dev libleptonica-dev

# CentOS、Red Hat
$ yum install -y tesseract

# 查看支持的语言：
$ tesseract --list-langs

# Ubuntu、Debian和Deepin 安装语言包
$ git clone https://github.com/tesseract-ocr/tessdata.gitsudo mv tessdata/* /usr/share/tesseract-ocr/tessdata

# 在CentOS和Red Hat系统下的迁移命令如下：
$ git clone https://github.com/tesseract-ocr/tessdata.gitsudo mv tessdata/* /usr/share/tesseract/tessdata
  
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

---

##### [post] /ocr/base64
> 传入base64图片进行识别

参数(from-data)
- base64img 被识别图片的base64
- lang 识别的词语库,默认是chi_sim

返回
```json
{
    "code": 0,
    "timestamp": 1636428522,
    "msg": "succeed",
    "data": "18778"
}
```

---

#### [get] /fanyi
> 翻译

参数
- word  需要翻译的词语

返回
```json
{"type":"ZH_CN2EN","errorCode":0,"elapsedTime":0,"translateResult":[[{"src":"你好","tgt":"hello"}]]}
```

---

#### [post] /pdf/txt
> 提取PDF,每页为一个文本

参数
- file  被提取的pdf文件

返回
```json
{
    "code": 0,
    "timestamp": 1636444276,
    "msg": "succeed",
    "data": [
        {
            "page": 1,
            "content": "- 65 -  附件2 普通高等学校本科专业目录 （2020年版） 说明： 1.本目录是在《普通高等学校本科专业目录（2012年）》基础上，增补近几年批准增设的目录外新专业而形成。 2.特设专业在专业代码后加T表示；国家控制布点专业在专业代码后加K表示。  序号 门类 专业类 专业代码 专业名称 学位授予门类 修业年限 增设年份 1 哲学 哲学类 010101 哲学 哲学 四年   2 哲学 哲学类 010102 逻辑学 哲学 四年   3 哲学 哲学类 010103K 宗教学 哲学 四年   4 哲学 哲学类 010104T 伦理学 哲学 四年   5 经济学 经济学类 020101 经济学 经济学 四年   6 经济学 经济学类 020102 经济统计学 经济学 四年   7 经济学 经济学类 020103T 国民经济管理 经济学 四年   8 经济学 经济学类 020104T 资源与环境经济学 经济学 四年   9 经济学 经济学类 020105T 商务经济学 经济学 四年   10 经济学 经济学类 020106T 能源经济 经济学 四年   11 经济学 经济学类 020107T 劳动经济学 经济学 四年 2016 12 经济学 经济学类 020108T 经济工程 经济学 四年 2017 13 经济学 经济学类 020109T 数字经济 经济学 四年 2018 14 经济学 财政学类 020201K 财政学 经济学 四年   15 经济学 财政学类 020202 税收学 经济学 四年   16 经济学 金融学类 020301K 金融学 经济学 四年   17 经济学 金融学类 020302 金融工程 经济学 四年   18 经济学 金融学类 020303 保险学 经济学 四年   19 经济学 金融学类 020304 投资学 经济学 四年   20 经济学 金融学类 020305T 金融数学 经济学 四年   21 经济学 金融学类 020306T 信用管理 管理学,经济学 四年   22 经济学 金融学类 020307T 经济与金融 经济学 四年   23 经济学 金融学类 020308T 精算学 理学,经济学 四年 2015 24 经济学 金融学类 020309T 互联网金融 经济学 四年 2016 "
        },
        {
            "page": 2,
            "content": "- 66 -  序号 .....
```

---

#### [post] /pdf/row
> 提取PDF,每页按行提取

参数
- file  被提取的pdf文件

返回
```json
{
    "code": 0,
    "timestamp": 1636444558,
    "msg": "succeed",
    "data": [
        {
            "page": 1,
            "content": [
                "附件2 ",
                "普通高等学校本科专业目录 ",
                "（2020年版） ",
                "说明： ",
                "1.（是在《普通高等学校本科专本目录业目录2012年）》基础上，增",
                "补近几年批准增设的目录外新专业而形成。 ",
                "2.特设专业在专业代码后加T表示；国家控制布点专业在专业代码后",
                "加K表示。 ",
                " ",
                "序学位授修业增设",
                "门类 专业类 专业代码 专业名称 ",
                "号 予门类 年限 年份 ",
                "1 哲学 哲学类 010101 哲学 哲学 四年   ",
                "2 哲学 哲学类 010102 逻辑学 哲学 四年   ",
                "3 哲学 哲学类 010103K 宗教学 哲学 四年   ",
                "4 哲学 哲学类 010104T 伦理学 哲学 四年   ",
                "5 经济学 经济学类 020101 经济学 经济学 四年   ",
                "6 经济学 经济学类 020102 学经济统计 经济学 四年   ",
                "7 经济学 经济学类 020103T 管理国民经济 经济学 四年   ",
                "资源与环境经济",
                "8 经济学 经济学类 020104T 经济学 四年   ",
                "学 ",
                "9 经济学 经济学类 020105T 学商务经济 经济学 四年   ",
                "10 经济学 经济学类 020106T 能源经济 经济学 四年   ",
                "11 经济学 经济学类 020107T 学劳动经济 经济学 四年 2016 ",
                "12 经济学 经济学类 020108T 经济工程 经济学 四年 2017 ",
                "13 经济学 经济学类 020109T 数字经济 经济学 四年 2018 ",
                "14 经济学 财政学类 020201K 财政学 经济学 四年   ",
                "15 经济学 财政学类 020202 税收学 经济学 四年   ",
                "16 经济学 金融学类 020301K 金融学 经济学 四年   ",
                "17 经济学 金融学类 020302 金融工程 经济学 四年   ",
                "18 经济学 金融学类 020303 保险学 经济学 四年   ",
                "19 经济学 金融学类 020304 投资学 经济学 四年   ",
                "20 经济学 金融学类 020305T 金融数学 经济学 四年   ",
                "管理学,",
                "21 经济学 金融学类 020306T 信用管理 四年   ",
                "经济学 ",
                "22 经济学 金融学类 020307T 融经济与金 经济学 四年   ",
                "理学,经",
                "23 经济学 金融学类 020308T 精算学 四年 2015 ",
                "济学 ",
                "24 经济学 金融学类 020309T 融互联网金 经济学 四年 2016 ",
                "- 65 - ",
                " "
            ]
        },
        {
            "page": 2,
            "content": [
                "序学位授修业增设",
                "门类 专业类  .....
```

---

#### [post] /pdf/table
> 提取PDF,每页提取标准表格数据, 注意: 只会提取表格数据

参数
- file  被提取的pdf文件

返回
```json
{
    "code": 0,
    "timestamp": 1636444640,
    "msg": "succeed",
    "data": [
        {
            "page": 1,
            "content": [
                {
                    "0": "序号 ",
                    "1": "门类 ",
                    "2": "专业类 ",
                    "3": "专业代码 ",
                    "4": "专业名称 ",
                    "5": "学位授予门类 ",
                    "6": "修业年限 ",
                    "7": "增设年份 "
                },
                {
                    "0": "1 ",
                    "1": "哲学 ",
                    "2": "哲学类 ",
                    "3": "010101 ",
                    "4": "哲学 ",
                    "5": "哲学 ",
                    "6": "四年 ",
                    "7": "  "
                },
                {
                    "0": "2 ",
                    "1": "哲学 ",
                    "2": "哲学类 ",
                    "3": "010102 ",
                    "4": "逻辑学 ",
                    "5": "哲学 ",
                    "6": "四年 ",
                    "7": "  "
                },
                {
                    "0": "3 ",
                    "1": "哲学 ",
                    "2": "哲学类 ",
                    "3": "010103K ",
                    "4": "宗教学 ",
                    "5": "哲学 ",
                    "6": "四年 ",
                    "7": "  "
                },
                {
                    "0": "4 ",
                    "1": "哲学 ",
                    "2": "哲学类 ",
                    "3": "010104T ",
                    "4": "伦理学 ",
                    "5": "哲学 ",
                    "6": "四年 ",
                    "7": "  "
                },......
```

---


## grpc 文档
> proto文件: https://github.com/mangenotwork/extras/api/WordHelper_Proto/wordhelper.proto

> 生成pb.go: https://github.com/mangenotwork/extras/script/wordhelper_pb.sh


## 编译
> 直接编译:  go build main.go

> 编译为docker: https://github.com/mangenotwork/extras/build/shortlink_build.sh
