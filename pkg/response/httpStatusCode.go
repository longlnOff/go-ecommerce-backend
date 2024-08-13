package response

const (
	ErrCodeSuccess		= 20000
	ErrCodeParamInvalid	= 20003
)

// message
var msg = map[int]string {
	ErrCodeSuccess: "success",
	ErrCodeParamInvalid: "email is invalid",
}