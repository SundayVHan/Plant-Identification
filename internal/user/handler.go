package user

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"plant_identification/internal/common"
)

func Register(c *gin.Context) {
	userName := c.Query("username")
	password := c.Query("password")

	token, err := RegisterAndIssueToken(userName, password)
	if err != nil {
		common.Error(c, err, common.ErrRegisterFailed)
		return
	}

	common.Success(c, "User successfully registered", gin.H{"token": token})
}

func Login(c *gin.Context) {
	userName := c.Query("username")
	password := c.Query("password")

	token, err := LoginAndIssueToken(userName, password)
	if err != nil {
		common.Error(c, err, http.StatusBadRequest)
		return
	}

	// 验证成功
	common.Success(c, "User successfully logged in", gin.H{"token": token})
}
