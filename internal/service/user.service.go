package service

import "github.com/longln/go-ecommerce-backend/internal/repo"

type UserService struct {
	userRepo *repo.UserRepo
}

func NewUserService() *UserService {
	return &UserService{
		userRepo: repo.NewUserRepo(),
	}
}

// user servuce
func (us *UserService) GetInforUserService() string {
	return us.userRepo.GetInforUser()
}