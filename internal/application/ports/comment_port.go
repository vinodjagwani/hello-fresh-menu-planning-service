package ports

import (
	"hello-fresh-menu-planning-service/internal/adaptors/rest/dto"
	"hello-fresh-menu-planning-service/internal/application/usecase"
)

type CommentCreatePort struct {
	CommentCreateUseCase *usecase.CommentCreateUseCase
}

func (p *CommentCreatePort) CreateComment(request *dto.CommentCreateRequest) (*dto.CommentCreateResponse, error) {
	return p.CommentCreateUseCase.CreateComment(request)
}
