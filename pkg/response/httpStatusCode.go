package response



const (
	ErrCodeSuccess 			= 20001	
	ErrCodeParamInvalid 	= 20003
	ErrInvalidToken			= 30001

	// for register code
	ErrCodeUserHasExists 	= 50001
)


var Msg = map[int]string {
	ErrCodeSuccess: "success",
	ErrCodeParamInvalid: "email is invalid",
	ErrInvalidToken: "token is invalid",

	// for register code
	ErrCodeUserHasExists: "user has exists",
}