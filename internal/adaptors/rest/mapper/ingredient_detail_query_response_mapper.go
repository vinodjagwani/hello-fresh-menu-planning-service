package mapper

import (
	"hello-fresh-menu-planning-service/internal/adaptors/rest/dto"
	"hello-fresh-menu-planning-service/internal/infra/repository/entity"
)

func MapIngredientToIngredientDetailQueryResponse(ingredient *entity.Ingredient) *dto.IngredientDetailQueryResponse {
	return &dto.IngredientDetailQueryResponse{IngredientId: ingredient.IngredientID.String(), IngredientName: ingredient.IngredientName, IngredientDescription: ingredient.IngredientDescription, IngredientUnit: ingredient.Unit}
}
