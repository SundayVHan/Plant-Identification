package article

import (
	"github.com/gin-gonic/gin"
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
