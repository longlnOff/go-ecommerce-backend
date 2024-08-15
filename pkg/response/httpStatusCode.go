package response

const (
	ErrCodeSuccess		= 20000
	ErrCodeParamInvalid	= 20003
	ErrCodeTokenInvalid	= 20004
)

// message
var msg = map[int]string {
	ErrCodeSuccess: "success",
	ErrCodeParamInvalid: "email is invalid",
	ErrCodeTokenInvalid: "token is invalid",
}