package user

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"plant_identification/internal/common"
	"strconv"
)

func Register(c *gin.Context) {
	userName := c.Query("username")
	password := c.Query("password")
	_kind := c.Query("kind")
	kind, _ := strconv.Atoi(_kind)

	token, err := RegisterAndIssueToken(userName, password, int64(kind))
	if err != nil {
		common.Error(c, err, common.ErrRegisterFailed)
		return
	}

	common.Success(c, "User successfully registered", gin.H{
		"token": token,
		"kind":  kind,
	})
}

func Login(c *gin.Context) {
	userName := c.Query("username")
	password := c.Query("password")

	token, kind, err := LoginAndIssueToken(userName, password)
	if err != nil {
		common.Error(c, err, http.StatusBadRequest)
		return
	}

	// 验证成功
	common.Success(c, "User successfully logged in", gin.H{
		"token": token,
		"kind":  kind,
	})
}
