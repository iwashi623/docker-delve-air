package repository

import "docker-delve-air/domain/model"

type UserRepository struct {
}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (r *UserRepository) GetUserByUserName(userName string) (*model.User, error) {
	return model.NewUser(userName), nil
}
