package user

type User struct {
	ID       int64  `json:"id" gorm:"primaryKey;autoIncrement"`
	Username string `json:"username" gorm:"unique;not null;size:255"`
	Password string `json:"password" gorm:"not null;size:255"`
	Kind     int64  `json:"kind" gorm:"not null"`
}
