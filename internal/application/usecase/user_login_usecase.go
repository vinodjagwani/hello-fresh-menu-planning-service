package usecase

import (
	"hello-fresh-menu-planning-service/internal/infra/repository"
	"hello-fresh-menu-planning-service/internal/infra/repository/entity"
)

type UserLoginUseCase struct {
	R *repository.UserRepository
}

type IUserLoginUseCase interface {
	LoginUser(email string, password string) (*entity.User, error)
}

func (uc *UserLoginUseCase) LoginUser(email string, password string) (*entity.User, error) {
	user, err := uc.R.FindUserByEmailAndPassword(email, password)
	return user, err
}
