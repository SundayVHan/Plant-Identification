package identification

import (
	"bytes"
	"errors"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/google/uuid"
	"gorm.io/gorm"
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

	url := "https://" + bucketName + "." + endpoint + "/" + objPath
	return url, err
}

func setHistory(userId int64, imageBase64 []byte, label string, response string) (History, error) {
	historyRecord := History{
		UserId:      userId,
		ImageBase64: imageBase64,
		Label:       label,
		Response:    response,
	}

	if err := database.DB.Create(&historyRecord).Error; err != nil {
		return historyRecord, err
	}

	return historyRecord, nil
}

func getHistoriesByUserId(userId int64) ([]History, error) {
	var histories []History

	// 使用 GORM 查询符合条件的记录
	if err := database.DB.Where("user_id = ?", userId).Find(&histories).Error; err != nil {
		return nil, err // 返回空切片和错误
	}

	return histories, nil // 返回查询到的历史记录和 nil 错误
}

func setStar(userId int64, historyId int64) error {
	history := History{
		ID: historyId,
	}

	// 查询指定历史记录
	if err := database.DB.First(&history).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("history record not found")
		}
		return err // 处理其他查询错误
	}

	// 确保是合法用户的记录
	if history.UserId != userId {
		return errors.New("user does not have permission to modify this history record")
	}

	// 更新星标状态
	history.Star = !history.Star // 切换星标状态

	// 保存更改
	if err := database.DB.Save(&history).Error; err != nil {
		return err // 处理保存错误
	}

	return nil // 返回 nil，表示成功
}
