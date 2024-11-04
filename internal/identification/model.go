package identification

import (
	"github.com/spf13/viper"
	"log"
	"time"
)

var (
	accessKeyId     string
	accessKeySecret string
	endpoint        string
	bucketName      string
	filePath        string
)


type QARequest struct {
	Image string `json:"image" binding:"required"`
}

type ReasonResponse struct {
	Label    string `json:"label"`
	Response string `json:"response"`
}

type History struct {
	ID          int64     `gorm:"primaryKey" json:"-"` // 主键自增
	UserId      int64     `gorm:"not null" json:"-"`
	ImageBase64 []byte    `gorm:"not null" json:"img_base64"` // 图片的Base64编码
	Label       string    `gorm:"not null" json:"label"`      // 标签
	Response    string    `gorm:"not null" json:"response"`   // 响应内容
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"time"` // 自动创建时间
}
