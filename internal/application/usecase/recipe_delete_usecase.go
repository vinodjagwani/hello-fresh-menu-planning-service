package usecase

import (
	"hello-fresh-menu-planning-service/internal/infra/repository"
)

type RecipeDeleteUseCase struct {
	R *repository.RecipeRepository
}

type IRecipeDeleteUseCase interface {
	DeleteRecipeById(recipeId string) error
}

func (uc *RecipeDeleteUseCase) DeleteRecipeById(recipeId string) error {
	err := uc.R.DeleteByRecipeId(recipeId)
	return err
}
