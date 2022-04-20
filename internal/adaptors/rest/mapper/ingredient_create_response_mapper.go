package mapper

import (
	"hello-fresh-menu-planning-service/internal/adaptors/rest/dto"
	"hello-fresh-menu-planning-service/internal/infra/repository/entity"
)

func MapIngredientToIngredientCreateResponse(ingredient *entity.Ingredient) *dto.IngredientCreateResponse {
	return &dto.IngredientCreateResponse{IngredientId: ingredient.IngredientID.String(), IngredientName: ingredient.IngredientName, IngredientDescription: ingredient.IngredientDescription, Unit: ingredient.Unit}
}
