package repository

import "mini-app-telegram/internal/domain"

type UserRepository struct {
	users map[int64]domain.User
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		users: make(map[int64]domain.User),
	}
}

func (u *UserRepository) CreateUser(user domain.User) {
	u.users[user.UserId] = user
}

func (u *UserRepository) GetUser(userId int64) (domain.User, bool) {
	user, ok := u.users[userId]
	if !ok {
		return domain.User{}, false
	}

	return user, true
}
