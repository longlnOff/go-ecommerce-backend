package service

import (
	"fmt"
	"strconv"
	"time"

	"github.com/longln/go-ecommerce-backend-api/global"
	"github.com/longln/go-ecommerce-backend-api/internal/utils/crypto"
	"github.com/longln/go-ecommerce-backend-api/internal/utils/sendto"
	"go.uber.org/zap"

	"github.com/longln/go-ecommerce-backend-api/internal/repo"
	"github.com/longln/go-ecommerce-backend-api/internal/utils/random"
	"github.com/longln/go-ecommerce-backend-api/pkg/response"
)

// INTERFACE_VERSION
// Chu y cach dat ten interface
type IUserService interface {
	Register(email string, purpose string) int
}

type userService struct {
	userRepo repo.IUserRepository
	userAuthRep repo.IUserAuthRepository

}

// Register implements IUserService.
func (us *userService) Register(email string, purpose string) int {
	// use redis to send OTP because it's faster than mongoDB

	// 0. hashEmail --> to protect information
	hashedEmail := crypto.GetHash(email)
	fmt.Printf("hash emal::%s", hashedEmail)

	// 5. check OTP is available

	// 6. handle user spam

	// 1. check email exists in db
	if us.userRepo.GetUserByEmail(email) {
		return response.ErrCodeUserHasExists
	}
	// 2. new OTP
	otp := random.GenerateSixDigitOTP()
	if purpose == "TEST_USER" {
		otp = 123456
	}

	fmt.Printf("OTP is :::%d\n", otp)

	// 3. Save OTP in Redis with expiration time
	err := us.userAuthRep.AddOTP(hashedEmail, otp, int64(10*time.Minute))
	if err != nil {
		return response.ErrInvalidOTP
	}
	// 4. Send email
	err = sendto.SendTemplateEmailOTP([]string{email}, "smtp.gmail.com", "otp-auth.html", map[string]interface{}{"otp": strconv.Itoa(otp)})
	if err != nil {
		global.Logger.Error("Failed to send email", zap.Error(err))
		return response.ErrSendEmailOTP
	}

	return response.ErrCodeSuccess
}

func NewUserService(userRepo repo.IUserRepository, userAuthRepo repo.IUserAuthRepository) IUserService {
	return &userService{
		userRepo: userRepo,
		userAuthRep: userAuthRepo,
	}
}
