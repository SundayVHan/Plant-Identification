package identification

import (
	"bytes"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/google/uuid"
)

func UploadImage(imageData []byte) (string, error) {

	// 创建OSSClient实例
	client, err := oss.New(endpoint, accessKeyId, accessKeySecret)
	if err != nil {
		return "", err
	}

	// 获取存储空间
	bucket, err := client.Bucket(bucketName)
	if err != nil {
		return "", err
	}

	randomFileName := uuid.New().String() + ".jpg"
	objPath := filePath + "/" + randomFileName

	// 上传本地文件到OSS
	err = bucket.PutObject(objPath, bytes.NewReader(imageData))
	if err != nil {
		return "", err
	}

	url := bucketName + "." + endpoint + "/" + objPath
	return url, err
}
