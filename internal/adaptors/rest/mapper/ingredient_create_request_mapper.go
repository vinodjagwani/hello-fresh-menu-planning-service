package mapper

import (
	"hello-fresh-menu-planning-service/internal/adaptors/rest/dto"
	"hello-fresh-menu-planning-service/internal/infra/repository/entity"
)

func MapIngredientCreateRequestToIngredient(request *dto.IngredientCreateRequest) *entity.Ingredient {
	return &entity.Ingredient{IngredientName: request.IngredientName, IngredientDescription: request.IngredientDescription, Unit: request.IngredientUnit}
}
