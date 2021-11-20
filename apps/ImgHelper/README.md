## ImgHelper 图片功能服务
> 生成二维码, 图片压缩， 图片水印， 生成文字图片， 图片固定剪切， 生成gif， 图片固定拼接， 图片基础信息获取，固定旋转， 格式转换;
>
> 提供 http/s, grpc api

## 配置说明
- app.runType 运行模式; dev:开发模式; prod:生产模式； test:测试模式；
- httpServer.open  是否启动http服务
- httpServer.prod  http服务端口
- grpcServer.open  是否启动grpc服务
- grpcServer.prod  grpc服务端口
- ttf              图片渲染的字体

```shell script
app:
  name: ImgHelper
  runType: dev

httpServer:
  open: true
  prod: 1261

grpcServer:
  open: true
  prod: 1262
  log: true

ttf: /ttf/Alibaba-PuHuiTi-Heavy.ttf
```

## Http接口文档

#### [get] /qrcode
> 生成二维码  QRCode

参数
- value 文本

返回
- 媒体文件


#### [get] /barcode
> 生成条形码  Barcode

参数
- value 文本

返回
- 媒体文件


#### [post] /image/info
> 图片信息获取

参数 from-data
- file 图片文件


返回
```json
{
  "code":0,
  "timestamp":1637226148,
  "msg":"succeed",
  "data":{
    "name":"微信图片_20210707110723.jpg",
    "size":47025,
    "type":"jpeg",
    "width":568,
    "height":779,
    "dpi":"568*779 dpi",
    "is_exif":false
  }
}
```

#### [post] /image/compress
> 图片压缩

参数 from-data
- file 图片文件
- level 压缩等级 

返回
- 媒体文件


#### [post] /watermark/txt
> 图片添加水印 - 文字水印

参数 from-data
- file 图片文件  必填
- txt 文字水印  必填
- color 文字水印颜色
- font_size  文字水印大小
- dpi  文字水印dpi
- position 文字水印位置 1:左下角, 2:居中, 3:左上角, 4:右上角, 5:右下角

返回
- 媒体文件

#### [post] /watermark/img
> 图片添加水印 - 图片水印

参数 from-data
- file 图片文件  必填
- logo 图片水印  必填

返回
- 媒体文件

#### [post] /watermark/logo
> 图片添加水印 - logo水印

参数 from-data
- file 图片文件  必填
- logo 图片水印  必填

返回
- 媒体文件

#### [post] /txt2img
> 生成文字图片, 应用场景: 文章转图片

参数 from-data
- txt 文本
- font_size 字大小
- dpi 文字dpi
- spacing 文字间距 默认2
- out_type 图片输出格式 png, jpg; 默认png

返回
- 媒体文件

#### [post] /img2gif
> 图片合成gif

参数 from-data
- file 图片文件 (multiple)

返回
- 媒体文件

#### [post] /img/revolve
> 图片旋转

参数 from-data
- file 图片文件 
- type 旋转类型; 90, 180, 270 

返回
- 媒体文件

#### [post] /img/center
> 图片居中

参数 from-data
- file 图片文件 

返回
- 媒体文件


#### [post] /img/stitching
> 图片拼接 垂直拼接

参数 from-data
- file_count 图片个数
- file_1 第一个图片文件 
- file_n 第n个图片文件

返回
- 媒体文件


#### [post] /img/sudoku
> 图片拼接 九宫格

参数 from-data
- file_count 图片个数
- file_1 第一个图片文件 
- file_n 第n个图片文件

返回
- 媒体文件


#### [post] /img/clipper/rect
> 图片剪裁 矩形裁剪

参数 from-data
- file 图片文件 
- x1 启始x坐标
- y1 启始y坐标
- x2 结束x坐标
- y2 结束y坐标

返回
- 媒体文件


## grpc 文档
> proto文件: https://github.com/mangenotwork/extras/api/ImgHelper_Proto/imghelper.proto

> 生成pb.go: https://github.com/mangenotwork/extras/script/imghelper_pb.sh


## 编译
> 直接编译:  go build main.go

> 编译为docker: https://github.com/mangenotwork/extras/build/imghelper_build.sh