package thumb

import (
	"github.com/gin-gonic/gin"
	"plant_identification/internal/common"
	user2 "plant_identification/internal/user"
)

func AddThumb(c *gin.Context) {
	_user, _ := c.Get("user")
	user, _ := _user.(user2.User)

	err := addThumb(user.ID)
	if err != nil {
		common.Error(c, err, common.ErrInternal)
		return
	}

	common.Success(c, "thumb success", nil)
}

func CountThumbs(c *gin.Context) {
	_user, _ := c.Get("user")
	user, _ := _user.(user2.User)

	if user.Kind == 0 {
		err := common.CustomError{
			Code:    common.ErrKindIsZero,
			Message: "only user whose kind = 1 can count thumbs",
		}
		common.Error(c, err, common.ErrUnauthorized)
		return
	}

	count, err := countThumbs()
	if err != nil {
		common.Error(c, err, common.ErrInternal)
		return
	}

	res := CountThumbsResponse{
		Sum: count,
	}
	common.Success(c, "count thumbs success", res)
}
