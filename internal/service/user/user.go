package user

import "mini-app-telegram/internal/domain"

type UserRepository interface {
	CreateUser(user domain.User)
	GetUser(userId int64) (domain.User, bool)
}

type UserService struct {
	userRepo UserRepository
}

func NewUserService(ur UserRepository) *UserService {
	return &UserService{
		userRepo: ur,
	}
}

func (u *UserService) GetUser(userId int64) domain.User {
	user, _ := u.userRepo.GetUser(userId)

	return user
}

func (u *UserService) UserExist(userId int64) bool {
	_, exist := u.userRepo.GetUser(userId)

	return exist
}

func (u *UserService) CreateUser(user domain.User) {
	u.userRepo.CreateUser(user)
}
