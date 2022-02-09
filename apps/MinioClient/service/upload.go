package service

import (
	"github.com/mangenotwork/extras/apps/MinioClient/model"
	"github.com/minio/minio-go/v6"
	"log"
	"mime/multipart"
)

func UploadFile(bucket string, file multipart.File, fileStat *multipart.FileHeader) (string, error) {
	log.Println("bucket = ", bucket)

	n, err := model.MinioClient.PutObject(bucket,
		fileStat.Filename, file, fileStat.Size,
		minio.PutObjectOptions{ContentType:fileStat.Header.Get("Content-Type")})
	if err != nil {
		return "", err
	}
	log.Println("Successfully uploaded bytes: ", n)
	return bucket+"/"+fileStat.Filename, nil
}