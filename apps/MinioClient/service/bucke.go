package service

import (
	"fmt"
	"github.com/mangenotwork/extras/apps/MinioClient/model"
	"log"
)

func MakeBucket(bucketName string) error {
	err := model.MinioClient.MakeBucket(bucketName, "location")
	if err != nil {
		// 检查存储桶是否已经存在。
		exists, err := model.MinioClient.BucketExists(bucketName)
		if err == nil && exists {
			return fmt.Errorf("存储桶 %s 已经存在", bucketName)
		} else {
			return err
		}
	}
	return nil
}

func BucketList() ([]string, error) {
	rse := make([]string, 0)
	lists, err := model.MinioClient.ListBuckets()
	if err != nil {
		return rse, err
	}

	for _, list := range lists {
		rse = append(rse, list.Name)
	}
	return rse, nil
}

func BucketFiles(bucketName string) []string {
	rse := make([]string, 0)
	objinfo := model.MinioClient.ListObjects(bucketName, "", true, make(chan struct{}))
	for obj := range objinfo {
		log.Println(obj.Key)
		rse = append(rse, bucketName+"/"+obj.Key)
	}
	return rse
}




