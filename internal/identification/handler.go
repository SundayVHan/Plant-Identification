package identification

import (
	"encoding/base64"
	"github.com/gin-gonic/gin"
	"plant_identification/internal/common"
	user2 "plant_identification/internal/user"
	"strconv"
)

func QA(c *gin.Context) {
	_user, _ := c.Get("user")
	user, _ := _user.(user2.User)

	var req QARequest
	if err := c.ShouldBindJSON(&req); err != nil {
		common.Error(c, err, common.ErrParamMissing)
		return
	}

	imageString := req.Image

	imageData, err := base64.StdEncoding.DecodeString(imageString)
	if err != nil {
		common.Error(c, err, common.ErrInternal)
		return
	}

	url, err := UploadImage(imageData)
	if err != nil {
		common.Error(c, err, common.ErrInternal)
		return
	}

	res, err := QueryLM(url, user.ID, imageData)
	if err != nil {
		common.Error(c, err, common.ErrInternal)
		return
	}

	common.Success(c, "identification success", res)
}

func GetHistory(c *gin.Context) {
	_user, _ := c.Get("user")
	user, _ := _user.(user2.User)

	histories, err := getHistoriesByUserId(user.ID)
	if err != nil {
		common.Error(c, err, common.ErrInternal)
		return
	}

	common.Success(c, "query histories success", histories)
}

func SetStar(c *gin.Context) {
	_user, _ := c.Get("user")
	user, _ := _user.(user2.User)

	_id := c.Query("id")
	id, _ := strconv.Atoi(_id)

	err := setStar(user.ID, int64(id))
	if err != nil {
		common.Error(c, err, common.ErrInternal)
		return
	}

	common.Success(c, "star success", nil)
}
