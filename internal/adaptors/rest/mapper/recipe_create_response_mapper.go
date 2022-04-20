package mapper

import (
	"hello-fresh-menu-planning-service/internal/adaptors/rest/dto"
	"hello-fresh-menu-planning-service/internal/infra/repository/entity"
)

func MapRecipeToRecipeCreateResponse(entity *entity.Recipe) *dto.RecipeCreateResponse {
	return &dto.RecipeCreateResponse{RecipeID: entity.RecipeID.String(), RecipeName: entity.RecipeName, RecipeDescription: entity.RecipeDescription, RecipeClassification: entity.RecipeClassification}
}
