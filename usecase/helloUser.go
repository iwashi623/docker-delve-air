package usecase

import "docker-delve-air/domain/repository"

type HelloUserUsecase struct {
	repo repository.UserRepository
}

func NewHelloUserUsecase(repo repository.UserRepository) *HelloUserUsecase {
	return &HelloUserUsecase{
		repo: repo,
	}
}

func (u *HelloUserUsecase) Exec(name string) (string, error) {
	user, err := u.repo.GetUserByUserName(name)
	if err != nil {
		return "", err
	}

	return "Hello, " + user.Name + "!", nil
}
