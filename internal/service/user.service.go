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
	userRepo *repo.IUserRepository
}

// Register implements IUserService.
func (us *userService) Register(email string, purpose string) int {

	// check email exists
	if us.userRepo.GetUserByEmail(email) {
		
	}
	return response.ErrCodeSuccess
}

func NewUserService(userRepo repo.IUserRepository) IUserService {
	return &userService{
		userRepo: &userRepo,
	}
}
