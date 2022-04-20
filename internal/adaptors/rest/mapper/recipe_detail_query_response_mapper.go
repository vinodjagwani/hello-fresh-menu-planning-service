package mapper

import (
	"hello-fresh-menu-planning-service/internal/adaptors/rest/dto"
	"hello-fresh-menu-planning-service/internal/infra/repository/entity"
)

func MapRecipeToRecipeDetailQueryResponse(entity *entity.Recipe) *dto.RecipeDetailQueryResponse {
	return &dto.RecipeDetailQueryResponse{RecipeId: entity.RecipeID.String(), RecipeName: entity.RecipeName, RecipeDescription: entity.RecipeDescription, RecipeClassification: entity.RecipeClassification}
}
