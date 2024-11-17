package article

import (
	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
	"os"
	"plant_identification/internal/database"
)

var client *openai.Client

func Init() {
	err := database.DB.AutoMigrate(&Article{})
	if err != nil {
		panic(err)
	}

	client = openai.NewClient(
		option.WithAPIKey(os.Getenv("HUNYUAN_API_KEY")),                 // 混元 APIKey
		option.WithBaseURL("https://api.hunyuan.cloud.tencent.com/v1/"), // 混元 endpoint
	)
}
