package ports

import (
	"hello-fresh-menu-planning-service/internal/adaptors/rest/dto"
	"hello-fresh-menu-planning-service/internal/application/usecase"
	"hello-fresh-menu-planning-service/internal/infra/repository/entity"
)

type UserLoginPort struct {
	UserLoginUseCase *usecase.UserLoginUseCase
}

type UserSignUpPort struct {
	UserSignUpUseCase *usecase.UserSignUpUseCase
}

func (p *UserLoginPort) LoginUser(email string, password string) (*entity.User, error) {
	return p.UserLoginUseCase.LoginUser(email, password)
}

func (p *UserSignUpPort) SignUp(request *dto.UserSignUpRequest) *dto.UserSignUpResponse {
	return p.UserSignUpUseCase.SignUp(request)
}
