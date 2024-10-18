package user

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"plant_identification/internal/router"
)

func Register(c *gin.Context) {
	userName := c.Query("username")
	password := c.Query("password")

	token, err := RegisterAndIssueToken(userName, password)
	if err != nil {
		router.Error(c, router.ErrRegisterFailed, err.Error(), http.StatusBadRequest)
		return
	}

	router.Success(c, "User successfully registered", gin.H{"token": token})
}

func Login(c *gin.Context) {
	userName := c.Query("username")
	password := c.Query("password")

	token, err := LoginAndIssueToken(userName, password)
	if err != nil {
		// 业务链异常
		router.Error(c, router.ErrLoginFailed, err.Error(), http.StatusBadRequest)
		return
	}

	// 验证成功
	router.Success(c, "User successfully logged in", gin.H{"token": token})
}
