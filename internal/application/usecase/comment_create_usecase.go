package usecase

import (
	"hello-fresh-menu-planning-service/internal/adaptors/rest/dto"
	"hello-fresh-menu-planning-service/internal/adaptors/rest/mapper"
	"hello-fresh-menu-planning-service/internal/infra/repository"
)

type CommentCreateUseCase struct {
	R *repository.CommentRepository
}

type ICommentCreateUseCase interface {
	CreateComment(request *dto.CommentCreateRequest) (*dto.CommentCreateResponse, error)
}

func (uc *CommentCreateUseCase) CreateComment(request *dto.CommentCreateRequest) (*dto.CommentCreateResponse, error) {
	var comment = mapper.MapCommentCreateRequestToComment(request)
	savedComment, err := uc.R.SaveComment(&comment)
	return mapper.MapCommentToCommentCreateResponse(savedComment), err
}
