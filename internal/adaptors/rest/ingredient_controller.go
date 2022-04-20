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
// @Summary Create ingredient
// @Schemes
// @Description body
// @Tags Ingredient
// @Accept json
// @Produce json
// @Param menuPlan body dto.IngredientCreateRequest true "Ingredient create request body"
// @Success 200 {object} dto.IngredientCreateResponse "Ingredient create response body"
// @Router /menu-planning-service/api/v1/ingredients [post]
func CreateIngredient(c *gin.Context) {
	var request dto.IngredientCreateRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	port := ports.IngredientCreatePort{IngredientCreateUseCase: &usecase.IngredientCreateUseCase{R: repository.GetIngredientRepository(postgres.GetDB())}}
	response, err := port.CreateIngredient(&request)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	c.JSON(http.StatusOK, &response)
}

// @BasePath /api/v1
// HelloFresh godoc
// @Summary Query ingredient by ingredient id
// @Schemes
// @Description body
// @Tags Ingredient
// @Accept json
// @Produce json
// @Param ingredientId  path  string  true "Ingredient ID"
// @Success 200 {object} dto.RecipeDetailQueryResponse "Ingredient detail query response body"
// @Router /menu-planning-service/api/v1/ingredients/{ingredientId} [get]
func FindIngredientById(c *gin.Context) {
	port := ports.IngredientQueryPort{IngredientQueryUseCase: &usecase.IngredientQueryUseCase{R: repository.GetIngredientRepository(postgres.GetDB())}}
	response, err := port.FindIngredientById(c.Param("ingredientId"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	c.JSON(http.StatusOK, &response)
}

// @BasePath /api/v1
// HelloFresh godoc
// @Summary Delete ingredient by ingredient id
// @Schemes
// @Description body
// @Tags Ingredient
// @Accept json
// @Produce json
// @Param ingredientId  path  string  true "Ingredient ID"
// @Success 204
// @Router /menu-planning-service/api/v1/ingredients/{ingredientId} [delete]
func DeleteIngredientById(c *gin.Context) {
	port := ports.IngredientDeletePort{IngredientPort: &usecase.IngredientDeleteUseCase{R: repository.GetIngredientRepository(postgres.GetDB())}}
	err := port.DeleteIngredientById(c.Param("ingredientId"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	c.JSON(http.StatusNoContent, "")
}
