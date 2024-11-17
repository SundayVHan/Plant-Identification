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
	ErrUsernameUsed
	ErrPasswordMismatch
	ErrUserNotRegistered
	ErrLMResponse
	ErrKindIsZero
)

type CustomError struct {
	Message string
	Code    WebCode
}

func (e CustomError) Error() string {
	return e.Message
}
