package repo

// type UserRepo struct {}

// func NewUserRepo() *UserRepo {
// 	return &UserRepo{}
// }

// // user repo
// func (ur *UserRepo) GetInforUser() string {
// 	return "longln"
// }

// Interface Version
type IUserRepository interface {
	GetUserByEmail(email string) bool
}

type userRepository struct {
}

// GetUserByEmail implements IUserRepository.
func (u *userRepository) GetUserByEmail(email string) bool {
	return true
}

func NewUserRepository() IUserRepository {
	return &userRepository{}
}
