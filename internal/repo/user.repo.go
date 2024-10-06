package repo

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
}

// GetUserByEmail implements IUserRepository.
func (ur *userRepository) GetUserByEmail(email string) bool {
	return true
}

func NewUserRepo() IUserRepository {
	return &userRepository{}
}
