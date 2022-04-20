package service

import (
	"hello-fresh-menu-planning-service/internal/adaptors/rest/dto"
	"hello-fresh-menu-planning-service/internal/adaptors/rest/mapper"
	"hello-fresh-menu-planning-service/internal/infra/postgres"
	"hello-fresh-menu-planning-service/internal/infra/repository"
	"hello-fresh-menu-planning-service/internal/infra/repository/entity"
)

func LoginUser(email string, password string) (*entity.User, error) {
	connection := repository.GetUserRepository(postgres.GetDB())
	user, err := connection.FindUserByEmailAndPassword(email, password)
	return user, err
}

func SignUp(request dto.UserSignUpRequest) *dto.UserSignUpResponse {
	var user = mapper.MapSignUpUserRequestToUser(&request)
	connection := repository.GetUserRepository(postgres.GetDB())
	saveUser, err := connection.SaveUser(&user)
	if err != nil {
		return &dto.UserSignUpResponse{}
	}
	return mapper.MapUserToSignUpUserResponse(saveUser)
}
