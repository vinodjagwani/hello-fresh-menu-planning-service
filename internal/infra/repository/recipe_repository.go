package repository

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
	"hello-fresh-menu-planning-service/internal/infra/repository/entity"
	"log"
)

type RecipeRepository struct {
	con *gorm.DB
}

func GetRecipeRepository(con *gorm.DB) *RecipeRepository {
	return &RecipeRepository{con: con}
}

type IRecipeRepository interface {
	SaveRecipe(recipe *entity.Recipe) (*entity.Recipe, error)
	UpdateRecipe(recipeId string, recipe *entity.Recipe) (*entity.Recipe, error)
	FindByRecipeId(recipeId string) (entity.Recipe, error)
}

func (r *RecipeRepository) SaveRecipe(recipe *entity.Recipe) (*entity.Recipe, error) {
	r.con.Create(&recipe)
	return recipe, r.con.Error
}

func (r *RecipeRepository) UpdateRecipe(recipeId string, recipe *entity.Recipe) (*entity.Recipe, error) {
	var updatedRecipe = entity.Recipe{}
	err := r.con.Where(&entity.Recipe{RecipeID: uuid.Must(uuid.FromString(recipeId))}).Updates(&recipe).First(&updatedRecipe).Error
	if err != nil {
		log.Printf("An error occurred while updating recipe with recipeId: %s with error %s", recipeId, err)
	}
	return &updatedRecipe, err
}

func (r *RecipeRepository) FindByRecipeId(recipeId string) (*entity.Recipe, error) {
	var result entity.Recipe
	err := r.con.Where(&entity.Recipe{RecipeID: uuid.Must(uuid.FromString(recipeId))}).First(&result).Error
	if err != nil {
		log.Printf("An error occurred while quering recipe with recipeId: %s with error %s", recipeId, err)
	}
	return &result, err
}

func (r *RecipeRepository) DeleteByRecipeId(recipeId string) error {
	var result entity.Recipe
	err := r.con.Where(&entity.Recipe{RecipeID: uuid.Must(uuid.FromString(recipeId))}).Delete(&result).Error
	if err != nil {
		log.Printf("An error occurred while deleting recipe with recipeId: %s with error %s", recipeId, err)
	}
	return err
}
