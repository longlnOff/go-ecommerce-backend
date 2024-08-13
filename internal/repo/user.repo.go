package repo

type UserRepo struct {}

func NewUserRepo() *UserRepo {
	return &UserRepo{}
}

// user repo
func (ur *UserRepo) GetInforUser() string {
	return "longln"
}