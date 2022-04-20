package usecase

import (
	"hello-fresh-menu-planning-service/internal/adaptors/rest/dto"
	"hello-fresh-menu-planning-service/internal/adaptors/rest/mapper"
	"hello-fresh-menu-planning-service/internal/infra/repository"
)

type RecipeQueryUseCase struct {
	R *repository.RecipeRepository
}

type IRecipeQueryUseCase interface {
	FindRecipeById(recipeId string) (*dto.RecipeDetailQueryResponse, error)
}

func (uc *RecipeQueryUseCase) FindRecipeById(recipeId string) (*dto.RecipeDetailQueryResponse, error) {
	recipe, err := uc.R.FindByRecipeId(recipeId)
	if err != nil {
		return &dto.RecipeDetailQueryResponse{}, err
	}
	return mapper.MapRecipeToRecipeDetailQueryResponse(recipe), err
}
