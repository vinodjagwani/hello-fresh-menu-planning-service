package usecase

import (
	"hello-fresh-menu-planning-service/internal/adaptors/rest/dto"
	"hello-fresh-menu-planning-service/internal/adaptors/rest/mapper"
	"hello-fresh-menu-planning-service/internal/infra/repository"
)

type RecipeCreateUseCase struct {
	R *repository.RecipeRepository
}

type IRecipeCreateUseCase interface {
	CreateRecipe(request *dto.RecipeCreateRequest) (dto.RecipeCreateResponse, error)
}

func (uc *RecipeCreateUseCase) CreateRecipe(request *dto.RecipeCreateRequest) (*dto.RecipeCreateResponse, error) {
	var recipe = mapper.MapRecipeCreateRequestToRecipe(request)
	savedRecipe, err := uc.R.SaveRecipe(&recipe)
	return mapper.MapRecipeToRecipeCreateResponse(savedRecipe), err
}
