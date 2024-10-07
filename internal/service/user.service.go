package service

import (
	"github.com/longln/go-ecommerce-backend-api/internal/repo"
	"github.com/longln/go-ecommerce-backend-api/pkg/response"
)

// INTERFACE_VERSION
// Chu y cach dat ten interface
type IUserService interface {
	Register(email string, purpose string) int
}

type userService struct {
	userRepo repo.IUserRepository
}

// Register implements IUserService.
func (us *userService) Register(email string, purpose string) int {
	// use redis to send OTP because it's faster than mongoDB

	// 0. hashEmail --> to protect information

	// check OTP is available

	// handle user spam

	// 1. check email exists in db

	// 2. new OTP

	// 3. Save OTP in Redis with expiration time

	// 4. Send email

	// check email exists
	if us.userRepo.GetUserByEmail(email) {
		return response.ErrCodeUserHasExists
	}
	return response.ErrCodeSuccess
}

func NewUserService(userRepo repo.IUserRepository) IUserService {
	return &userService{
		userRepo: userRepo,
	}
}
