package user

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Kind     int64  `json:"kind"`
}
