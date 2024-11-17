package identification

import (
	"encoding/json"
	"github.com/go-resty/resty/v2"
	"net/http"
	"plant_identification/internal/common"
	"plant_identification/internal/database"
)

func Init() {
	err := database.DB.AutoMigrate(&History{})
	if err != nil {
		panic(err)
	}
}

func QueryLM(imgUrl string, userId int64, imageData []byte) (res ReasonResponse, retErr error) {
	client := resty.New()

	resp, err := client.R().
		SetHeader("Accept", "application/json").
		SetBody(map[string]interface{}{
			"img_url": imgUrl,
		}).
		Post("http://127.0.0.1:5000/get_response")

	if err != nil || resp.StatusCode() != http.StatusOK {
		retErr = common.CustomError{
			Code:    common.ErrLMResponse,
			Message: "Language Model response failed",
		}
		return
	}

	if err = json.Unmarshal(resp.Body(), &res); err != nil {
		retErr = err
		return
	}

	history, err := setHistory(userId, imageData, res.Label, res.Response)
	if err != nil {
		retErr = err
		return
	}

	res.ID = history.ID

	return
}
