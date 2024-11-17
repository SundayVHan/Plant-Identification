package thumb

type Thumb struct {
	UserId int64 `gorm:"not null" json:"-"`
}

type CountThumbsResponse struct {
	Sum int64 `json:"sum"`
}
