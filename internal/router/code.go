package router

type WebCode int

const (
	ErrUnauthorized WebCode = iota + 10000
	ErrTokenInvalid
	ErrUserNotFound
	ErrRegisterFailed
	ErrLoginFailed
)
