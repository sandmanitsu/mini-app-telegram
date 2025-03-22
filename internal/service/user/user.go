package user

import (
	"fmt"
	"mini-app-telegram/internal/domain"
)

type UserRepository interface {
	CreateUser(user domain.User) error
	GetUser(userId int64) (domain.User, error)
}

type UserService struct {
	userRepo UserRepository
}

func NewUserService(ur UserRepository) *UserService {
	return &UserService{
		userRepo: ur,
	}
}

func (u *UserService) GetUser(userId int64) (domain.User, error) {
	return u.userRepo.GetUser(userId)
}

func (u *UserService) CreateUser(user domain.User) error {
	return u.userRepo.CreateUser(user)
}

func (u *UserService) UserExist(userId int64) bool {
	_, err := u.userRepo.GetUser(userId)
	fmt.Println(err)
	return err == nil
}
