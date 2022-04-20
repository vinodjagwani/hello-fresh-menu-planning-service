package mapper

import (
	"hello-fresh-menu-planning-service/internal/adaptors/rest/dto"
	"hello-fresh-menu-planning-service/internal/infra/repository/entity"
)

func MapCommentCreateRequestToComment(request *dto.CommentCreateRequest) entity.Comment {
	return entity.Comment{Comment: request.Comment, UserID: request.UserID, MenuPlanID: request.MenuPlanId}
}
