package service

import (
	"bytes"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/mangenotwork/extras/apps/MinioClient/model"
	"github.com/mangenotwork/extras/common/httpser"
	"github.com/mangenotwork/extras/common/logger"
	"github.com/mangenotwork/extras/common/utils"
	"github.com/minio/minio-go/v6"
	"github.com/nfnt/resize"
)

// GetFile 请求文件的参数介绍
// compact 压缩等级
//
func GetFile(w http.ResponseWriter, bucket, obj, compact string) {
	bucket = strings.TrimPrefix(bucket, "/")
	log.Println("bucket = ", bucket)
	log.Println("obj = ", obj)
	object, err := model.MinioClient.GetObject(bucket, obj, minio.GetObjectOptions{})
	if err != nil {
		log.Println(err)
		return
	}
	defer object.Close()

	objectInfo, err :=  object.Stat()
	logger.Info(objectInfo, err)

	// 图片压缩
	if strings.Index(objectInfo.ContentType, "jpeg") != -1 || strings.Index(objectInfo.ContentType, "gif") != -1 ||
		strings.Index(objectInfo.ContentType, "png") != -1  || strings.Index(objectInfo.ContentType, "bmp") != -1 ||
		strings.Index(objectInfo.ContentType, "wbmp") != -1 {
		img, str, err := image.Decode(object)
		if err != nil {
			httpser.OutErrBody(w, 2001, err)
			return
		}
		b := img.Bounds()
		width := b.Max.X
		levelInt := utils.Str2Int(compact)
		if levelInt <= 0 {
			levelInt = 1
		}
		out := ImgCompress(img, width/levelInt, 0, str)
		_,_=w.Write(out)
		return
	}

	w.WriteHeader(http.StatusOK)
	_, err = io.Copy(w, object)
	if err != nil {
		log.Println("err = ", err)
	}
}


// Nearest-neighbor interpolation
//	NearestNeighbor InterpolationFunction = iota
	// Bilinear interpolation
//	Bilinear
	// Bicubic interpolation (with cubic hermite spline)
//	Bicubic
	// Mitchell-Netravali interpolation
//	MitchellNetravali
	// Lanczos interpolation (a=2)
//	Lanczos2
	// Lanczos interpolation (a=3)
//	Lanczos3

// 图片压缩
func ImgCompress(img image.Image, width, height int, outType string) []byte {
	// resize.Resize 使用插值函数interp创建具有新尺寸（宽度，高度）的缩放图像。 如果宽度或高度设置为0，则将其设置为保留宽高比值。
	m := resize.Resize(uint(width), uint(height), img, resize.Lanczos3)
	out := new(bytes.Buffer)

	switch outType {
	case "png","PNG":
		_=png.Encode(out, m)
	case "jpg", "jpeg", "JPG", "JPEG":
		_=jpeg.Encode(out, m, nil)
	case "gif", "GIF":
		_=gif.Encode(out, m, nil)
	}
	return out.Bytes()
}

