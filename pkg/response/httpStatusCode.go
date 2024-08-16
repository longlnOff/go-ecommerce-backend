package response

const (
	ErrCodeSuccess		= 20000
	ErrCodeParamInvalid	= 20003
	ErrCodeTokenInvalid	= 20004
	// Register email
	ErrCodeUserHasExisted	= 50001	// user has existed

)

// message
var msg = map[int]string {
	ErrCodeSuccess: "success",
	ErrCodeParamInvalid: "email is invalid",
	ErrCodeTokenInvalid: "token is invalid",

	ErrCodeUserHasExisted: "user has existed",
}