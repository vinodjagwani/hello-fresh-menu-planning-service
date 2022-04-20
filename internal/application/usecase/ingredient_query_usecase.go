package usecase

import (
	"hello-fresh-menu-planning-service/internal/adaptors/rest/dto"
	"hello-fresh-menu-planning-service/internal/adaptors/rest/mapper"
	"hello-fresh-menu-planning-service/internal/infra/repository"
)

type IngredientQueryUseCase struct {
	R *repository.IngredientRepository
}

type IIngredientQueryUseCase interface {
	FindIngredientById(ingredientId string) (*dto.IngredientDetailQueryResponse, error)
}

func (uc *IngredientQueryUseCase) FindIngredientById(ingredientId string) (*dto.IngredientDetailQueryResponse, error) {
	ingredient, err := uc.R.FindByIngredientId(ingredientId)
	return mapper.MapIngredientToIngredientDetailQueryResponse(ingredient), err
}
