package identification

import (
	"encoding/base64"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
	"net/http"
	"plant_identification/internal/common"
)

type QARequest struct {
	Image string `json:"image" binding:"required"`
}

type ReasonResponse struct {
	Label     string `json:"label"`
	Responese string `json:"responese"`
}

func QA(c *gin.Context) {
	var req QARequest
	if err := c.ShouldBindJSON(&req); err != nil {
		common.Error(c, common.ErrParamMissing, "json param [image] missing", http.StatusBadRequest)
		return
	}

	imageString := req.Image

	imageData, err := base64.StdEncoding.DecodeString(imageString)
	if err != nil {
		common.Error(c, common.ErrInternal, "base64 decode image failed", http.StatusInternalServerError)
		return
	}

	url, err := UploadImage(imageData)
	if err != nil {
		common.Error(c, common.ErrInternal, "upload image failed", http.StatusInternalServerError)
		return
	}

	client := resty.New()

	resp, err := client.R().
		SetHeader("Accept", "application/json").
		SetBody(map[string]interface{}{
			"img_url": url,
		}).
		Get("http://localhost:5000/get_response")

	if err != nil || resp.StatusCode() != http.StatusOK {
		common.Error(c, common.ErrInternal, "identification process failed", http.StatusInternalServerError)
		return
	}

	var res ReasonResponse
	if err = json.Unmarshal(resp.Body(), &res); err != nil {
		common.Error(c, common.ErrInternal, "identification process failed", http.StatusInternalServerError)
		return
	}

	common.Success(c, "identification success", res)
}
