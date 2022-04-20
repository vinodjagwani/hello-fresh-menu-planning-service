package usecase

import (
	"hello-fresh-menu-planning-service/internal/adaptors/rest/dto"
	"hello-fresh-menu-planning-service/internal/adaptors/rest/mapper"
	"hello-fresh-menu-planning-service/internal/infra/repository"
)

type UserSignUpUseCase struct {
	R *repository.UserRepository
}

type IUserSignUpUseCase interface {
	SignUp(request *dto.UserSignUpRequest) *dto.UserSignUpResponse
}

func (uc *UserSignUpUseCase) SignUp(request *dto.UserSignUpRequest) *dto.UserSignUpResponse {
	var user = mapper.MapSignUpUserRequestToUser(request)
	saveUser, err := uc.R.SaveUser(&user)
	if err != nil {
		return &dto.UserSignUpResponse{}
	}
	return mapper.MapUserToSignUpUserResponse(saveUser)
}
