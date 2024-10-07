package repo

import (
	"context"

	"github.com/longln/go-ecommerce-backend-api/global"
	"github.com/longln/go-ecommerce-backend-api/internal/database"
)

// type UserRepo struct {

// }

// func NewUserRepo() *UserRepo {
// 	return &UserRepo{}
// }

// func (ur *UserRepo) GetUserInfoRepo() string {
// 	return "longln"
// }

type IUserRepository interface {
	GetUserByEmail(email string) bool
}

type userRepository struct {
	sqlc *database.Queries
}

// GetUserByEmail implements IUserRepository.
func (ur *userRepository) GetUserByEmail(email string) bool {
	user, err := ur.sqlc.GetUserByEmailSQLC(context.Background(), email)
	if err != nil {
		return false
	}
	return user.UsrID != 0
}

func NewUserRepository() IUserRepository {
	return &userRepository{
		sqlc: database.New(global.Mdbc),
	}
}
