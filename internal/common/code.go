package common

type WebCode int

const (
	ErrUnauthorized WebCode = iota + 10000
	ErrTokenInvalid
	ErrUserNotFound
	ErrRegisterFailed
	ErrLoginFailed
	ErrParamMissing
	ErrInternal
)
