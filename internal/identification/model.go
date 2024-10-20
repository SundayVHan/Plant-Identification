package identification

import (
	"github.com/spf13/viper"
	"log"
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
