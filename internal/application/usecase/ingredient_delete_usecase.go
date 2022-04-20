package usecase

import (
	"hello-fresh-menu-planning-service/internal/infra/repository"
)

type IngredientDeleteUseCase struct {
	R *repository.IngredientRepository
}

type IIngredientDeleteUseCase interface {
	DeleteIngredientById(ingredientId string) error
}

func (uc *IngredientDeleteUseCase) DeleteIngredientById(ingredientId string) error {
	return uc.R.DeleteByIngredientId(ingredientId)
}
