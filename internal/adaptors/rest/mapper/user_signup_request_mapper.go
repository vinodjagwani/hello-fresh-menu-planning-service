package mapper

import (
	"hello-fresh-menu-planning-service/internal/adaptors/rest/dto"
	"hello-fresh-menu-planning-service/internal/infra/repository/entity"
)

func MapSignUpUserRequestToUser(request *dto.UserSignUpRequest) entity.User {
	return entity.User{UserName: request.UserName, Email: request.Email, UserType: request.UserType, Password: request.Password}
}
