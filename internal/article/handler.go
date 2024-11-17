package article

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
	"plant_identification/internal/common"
	user2 "plant_identification/internal/user"
)

func PublishArticle(c *gin.Context) {
	_user, _ := c.Get("user")
	user, _ := _user.(user2.User)

	if user.Kind == 0 {
		err := common.CustomError{
			Code:    common.ErrKindIsZero,
			Message: "only user whose kind = 1 can publish article",
		}
		common.Error(c, err, common.ErrUnauthorized)
		return
	}

	var req PublishArticleRequest
	if err := c.ShouldBind(&req); err != nil {
		common.Error(c, err, common.ErrParamMissing)
		return
	}

	err := addArticle(req.Text, req.ImgBase64, req.Title)
	if err != nil {
		common.Error(c, err, common.ErrInternal)
		return
	}

	common.Success(c, "publish article success", nil)
}

func FetchArticle(c *gin.Context) {
	articles, err := getArticles()
	if err != nil {
		common.Error(c, err, common.ErrInternal)
		return
	}

	common.Success(c, "fetch articles success", articles)
}

func GenerateArticle(c *gin.Context) {
	var req GenerateArticleRequest
	if err := c.Bind(&req); err != nil {
		common.Error(c, err, common.ErrParamMissing)
		return
	}

	ctx := context.Background()

	completion, err := client.Chat.Completions.New(ctx,
		openai.ChatCompletionNewParams{
			Messages: openai.F([]openai.ChatCompletionMessageParamUnion{
				openai.SystemMessage("根据用户的输入，帮助用户续写一段一篇短文"),
				openai.UserMessage(req.Text),
			}),
			Model: openai.F("hunyuan-pro"),
		},
		option.WithJSONSet("enable_enhancement", true), // <- 自定义参数
	)

	if err != nil {
		common.Error(c, err, common.ErrInternal)
		return
	}

	res := GenerateArticleResponse{
		Generation: completion.Choices[0].Message.Content,
	}

	common.Success(c, "generate article success", res)
}
