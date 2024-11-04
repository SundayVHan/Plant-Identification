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

func init() {
	// 设置 viper 读取配置文件的名称和格式
	viper.SetConfigName("config") // 指定配置文件名称（不含扩展名）
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config") // 假设 config.yaml 在项目根目录

	// 读取配置文件
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	// 解析配置项
	accessKeyId = viper.GetString("aliyun.accessKeyId")
	accessKeySecret = viper.GetString("aliyun.accessKeySecret")
	endpoint = viper.GetString("aliyun.endpoint")
	bucketName = viper.GetString("aliyun.bucketName")
	filePath = viper.GetString("aliyun.filePath")
}

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
