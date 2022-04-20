package rest

import (
	"github.com/gin-gonic/gin"
	"hello-fresh-menu-planning-service/internal/adaptors/rest/dto"
	"hello-fresh-menu-planning-service/internal/application/ports"
	"hello-fresh-menu-planning-service/internal/application/usecase"
	"hello-fresh-menu-planning-service/internal/infra/postgres"
	"hello-fresh-menu-planning-service/internal/infra/repository"
	"net/http"
)

// @BasePath /api/v1
// HelloFresh godoc
// @Summary Create recipe
// @Schemes
// @Description body
// @Tags Recipe
// @Accept json
// @Produce json
// @Param recipeCreateRequest body dto.RecipeCreateRequest true "Recipe create request body"
// @Success 200 {object} dto.RecipeCreateResponse "Recipe create response body"
// @Router /menu-planning-service/api/v1/recipes [post]
func CreateRecipe(c *gin.Context) {
	var request dto.RecipeCreateRequest
	if err := c.BindJSON(&request); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	port := ports.RecipeCreatePort{RecipeCreateUseCase: &usecase.RecipeCreateUseCase{R: repository.GetRecipeRepository(postgres.GetDB())}}
	response, err := port.CreateRecipe(&request)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, &response)
}

// @BasePath /api/v1
// HelloFresh godoc
// @Summary Update recipe
// @Schemes
// @Description body
// @Tags Recipe
// @Accept json
// @Produce json
// @Param recipeId  path  string  true "Recipe ID"
// @Param recipeUpdateRequestBody body dto.RecipeCreateRequest true "Recipe update request body"
// @Success 200 {object} dto.RecipeUpdateResponse "Recipe update response body"
// @Router /menu-planning-service/api/v1/recipes/{recipeId} [put]
func UpdateRecipe(c *gin.Context) {
	var request dto.RecipeUpdateRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	port := ports.RecipeUpdatePort{RecipeUpdateUseCase: &usecase.RecipeUpdateUseCase{R: repository.GetRecipeRepository(postgres.GetDB())}}
	response, err := port.UpdateRecipe(c.Param("recipeId"), &request)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, &response)
}

// @BasePath /api/v1
// HelloFresh godoc
// @Summary Query recipe by recipe id
// @Schemes
// @Description body
// @Tags Recipe
// @Accept json
// @Produce json
// @Param recipeId  path  string  true "Recipe ID"
// @Success 200 {object} dto.RecipeDetailQueryResponse "Recipe detail query response body"
// @Router /menu-planning-service/api/v1/recipes/{recipeId} [get]
func FindRecipeById(c *gin.Context) {
	port := ports.RecipeQueryPort{RecipeQueryUseCase: &usecase.RecipeQueryUseCase{R: repository.GetRecipeRepository(postgres.GetDB())}}
	response, err := port.FindRecipeById(c.Param("recipeId"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, &response)
}

// @BasePath /api/v1
// HelloFresh godoc
// @Summary Delete recipe by recipe id
// @Schemes
// @Description body
// @Tags Recipe
// @Accept json
// @Produce json
// @Param recipeId  path  string  true "Recipe ID"
// @Success 204
// @Router /menu-planning-service/api/v1/recipes/{recipeId} [delete]
func DeleteRecipeById(c *gin.Context) {
	port := ports.RecipeDeletePort{RecipeDeleteUseCase: &usecase.RecipeDeleteUseCase{R: repository.GetRecipeRepository(postgres.GetDB())}}
	err := port.DeleteRecipeById(c.Param("recipeId"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusNoContent, "")
}
