package common

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code    WebCode     `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func Success(c *gin.Context, message string, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Code:    0,
		Message: message,
		Data:    data,
	})
}

func Error(c *gin.Context, err error, unexpectedCode WebCode) {
	var cErr *CustomError
	if errors.As(err, &cErr) {
		c.JSON(http.StatusBadRequest, Response{
			Code:    cErr.Code,
			Message: cErr.Message,
			Data:    nil,
		})
	} else {
		c.JSON(http.StatusInternalServerError, Response{
			Code:    unexpectedCode,
			Message: err.Error(),
			Data:    nil,
		})
	}
}

func Abort(c *gin.Context, code WebCode, message string, httpStatusCode int) {
	c.AbortWithStatusJSON(httpStatusCode, Response{
		Code:    code,
		Message: message,
		Data:    nil,
	})
}
