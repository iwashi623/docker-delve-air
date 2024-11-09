package repository

import "docker-delve-air/domain/model"

type UserRepository interface {
	GetUserByUserName(userName string) (*model.User, error)
}
