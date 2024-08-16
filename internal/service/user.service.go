package service

import (
	"github.com/longln/go-ecommerce-backend/internal/repo"
	"github.com/longln/go-ecommerce-backend/pkg/response"
)

// type UserService struct {
// 	userRepo *repo.UserRepo
// }

// func NewUserService() *UserService {
// 	return &UserService{
// 		userRepo: repo.NewUserRepo(),
// 	}
// }

// // user servuce
// func (us *UserService) GetInforUserService() string {
// 	return us.userRepo.GetInforUser()
// }

// Interface Version
type IUserService interface {
	Register(email string, purpose string) int
}

type userService struct {
	userRepo repo.IUserRepository
	// ... other fields
}

// Register implements IUserService.
func (us *userService) Register(email string, purpose string) int {
	// 1. Check email exists
	if us.userRepo.GetUserByEmail(email) {
		return response.ErrCodeUserHasExisted
	}
	return response.ErrCodeSuccess
}

func NewUserService(userRepo repo.IUserRepository) IUserService {
	return &userService{
		userRepo: userRepo,
	}
}
