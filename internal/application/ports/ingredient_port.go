package ports

import (
	"hello-fresh-menu-planning-service/internal/adaptors/rest/dto"
	"hello-fresh-menu-planning-service/internal/application/usecase"
)

type IngredientCreatePort struct {
	IngredientCreateUseCase *usecase.IngredientCreateUseCase
}

type IngredientQueryPort struct {
	IngredientQueryUseCase *usecase.IngredientQueryUseCase
}

type IngredientDeletePort struct {
	IngredientPort *usecase.IngredientDeleteUseCase
}

func (p *IngredientCreatePort) CreateIngredient(request *dto.IngredientCreateRequest) (*dto.IngredientCreateResponse, error) {
	return p.IngredientCreateUseCase.CreateIngredient(request)
}

func (p *IngredientQueryPort) FindIngredientById(ingredientId string) (*dto.IngredientDetailQueryResponse, error) {
	return p.IngredientQueryUseCase.FindIngredientById(ingredientId)
}

func (p *IngredientDeletePort) DeleteIngredientById(ingredientId string) error {
	return p.IngredientPort.DeleteIngredientById(ingredientId)
}
