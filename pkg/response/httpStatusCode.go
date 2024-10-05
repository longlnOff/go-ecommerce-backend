package response



const (
	ErrCodeSuccess 			= 20001	
	ErrCodeParamInvalid 	= 20003
	ErrInvalidToken			= 30001
)


var Msg = map[int]string {
	ErrCodeSuccess: "success",
	ErrCodeParamInvalid: "email is invalid",
	ErrInvalidToken: "token is invalid",
}