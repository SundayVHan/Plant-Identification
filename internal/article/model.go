package article

type Article struct {
	ID        int64  `gorm:"primary_key;auto_increment" json:"-"`
	Text      string `json:"text"`
	ImgBase64 string `json:"img_base64"`
	Title     string `json:"title"`
}

type PublishArticleRequest struct {
	Text      string `json:"text"`
	ImgBase64 string `json:"img_base64"`
	Title     string `json:"title"`
}
