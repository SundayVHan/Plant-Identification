package identification

import (
	"bytes"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/google/uuid"
	"plant_identification/internal/database"
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

func setHistory(userId int64, imageBase64 []byte, label string, response string) error {
	historyRecord := History{
		UserId:      userId,
		ImageBase64: imageBase64,
		Label:       label,
		Response:    response,
	}

	if err := database.DB.Create(&historyRecord).Error; err != nil {
		return err
	}

	return nil
}

func getHistoriesByUserId(userId int64) ([]History, error) {
	var histories []History

	// 使用 GORM 查询符合条件的记录
	if err := database.DB.Where("user_id = ?", userId).Find(&histories).Error; err != nil {
		return nil, err // 返回空切片和错误
	}

	return histories, nil // 返回查询到的历史记录和 nil 错误
}
