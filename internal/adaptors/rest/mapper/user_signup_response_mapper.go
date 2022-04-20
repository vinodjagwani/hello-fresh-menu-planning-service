package mapper

import (
	"hello-fresh-menu-planning-service/internal/adaptors/rest/dto"
	"hello-fresh-menu-planning-service/internal/infra/repository/entity"
)

func MapUserToSignUpUserResponse(user *entity.User) *dto.UserSignUpResponse {
	return &dto.UserSignUpResponse{UserId: user.UserID.String()}
}
