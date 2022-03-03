package service

import (
	"github.com/mangenotwork/extras/apps/MinioClient/model"
	"github.com/mangenotwork/extras/common/utils"
	"github.com/minio/minio-go/v6"
	"log"
	"mime/multipart"
	"path"
	"time"
)

func UploadFile(bucket string, file multipart.File, fileStat *multipart.FileHeader) (string, error) {
	log.Println("bucket = ", bucket)

	newFileName := utils.MD5(fileStat.Filename+time.Now().String())+ path.Ext(fileStat.Filename)
	n, err := model.MinioClient.PutObject(bucket,
		newFileName, file, fileStat.Size,
		minio.PutObjectOptions{ContentType:fileStat.Header.Get("Content-Type")})
	if err != nil {
		return "", err
	}
	log.Println("Successfully uploaded bytes: ", n)
	return bucket+"/"+newFileName, nil
}