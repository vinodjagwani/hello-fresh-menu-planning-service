package usecase

import (
	"hello-fresh-menu-planning-service/internal/adaptors/rest/dto"
	"hello-fresh-menu-planning-service/internal/adaptors/rest/mapper"
	"hello-fresh-menu-planning-service/internal/infra/repository"
)

type RecipeUpdateUseCase struct {
	R *repository.RecipeRepository
}

type IRecipeUpdateUseCase interface {
	UpdateRecipe(recipeId string, request *dto.RecipeUpdateRequest) (dto.RecipeUpdateResponse, error)
}

func (uc *RecipeUpdateUseCase) UpdateRecipe(recipeId string, request *dto.RecipeUpdateRequest) (*dto.RecipeUpdateResponse, error) {
	var recipe = mapper.MapRecipeUpdateRequestToRecipe(request)
	savedRecipe, err := uc.R.UpdateRecipe(recipeId, &recipe)
	return mapper.MapRecipeToRecipeUpdateResponse(savedRecipe), err
}
