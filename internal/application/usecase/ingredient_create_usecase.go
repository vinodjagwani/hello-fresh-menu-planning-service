package usecase

import (
	"hello-fresh-menu-planning-service/internal/adaptors/rest/dto"
	"hello-fresh-menu-planning-service/internal/adaptors/rest/mapper"
	"hello-fresh-menu-planning-service/internal/infra/repository"
)

type IngredientCreateUseCase struct {
	R *repository.IngredientRepository
}

type IIngredientCreateUseCase interface {
	CreateIngredient(request *dto.IngredientCreateRequest) (*dto.IngredientCreateResponse, error)
}

func (uc *IngredientCreateUseCase) CreateIngredient(request *dto.IngredientCreateRequest) (*dto.IngredientCreateResponse, error) {
	var ingredient = mapper.MapIngredientCreateRequestToIngredient(request)
	savedIngredient, err := uc.R.SaveIngredient(ingredient)
	return mapper.MapIngredientToIngredientCreateResponse(savedIngredient), err
}
