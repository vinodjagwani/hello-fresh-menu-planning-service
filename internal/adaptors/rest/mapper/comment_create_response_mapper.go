package mapper

import (
	"hello-fresh-menu-planning-service/internal/adaptors/rest/dto"
	"hello-fresh-menu-planning-service/internal/infra/repository/entity"
)

func MapCommentToCommentCreateResponse(comment *entity.Comment) *dto.CommentCreateResponse {
	return &dto.CommentCreateResponse{CommentId: comment.CommentID.String()}
}
