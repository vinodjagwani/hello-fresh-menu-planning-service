package repository

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
	"hello-fresh-menu-planning-service/internal/infra/repository/entity"
	"log"
)

type IngredientRepository struct {
	con *gorm.DB
}

func GetIngredientRepository(con *gorm.DB) *IngredientRepository {
	return &IngredientRepository{con: con}
}

type IIngredientRepository interface {
	SaveIngredient(ingredient *entity.Ingredient) (*entity.Ingredient, error)
	FindByIngredientId(ingredient string) (*entity.Ingredient, error)
	DeleteByIngredientId(ingredientId string) error
}

func (r *IngredientRepository) SaveIngredient(ingredient *entity.Ingredient) (*entity.Ingredient, error) {
	result := r.con.Create(&ingredient)
	return ingredient, result.Error
}

func (r *IngredientRepository) FindByIngredientId(ingredient string) (*entity.Ingredient, error) {
	var result entity.Ingredient
	err := r.con.Where(&entity.Ingredient{IngredientID: uuid.Must(uuid.FromString(ingredient))}).Find(&result).Error
	if err != nil {
		log.Printf("An error occurred while quering ingredient with ingredient: %s with error %s", ingredient, err)
	}
	return &result, err
}

func (r *IngredientRepository) DeleteByIngredientId(ingredientId string) error {
	var result entity.Ingredient
	err := r.con.Where(&entity.Ingredient{IngredientID: uuid.Must(uuid.FromString(ingredientId))}).Delete(&result).Error
	if err != nil {
		log.Printf("An error occurred while deleting ingredient with ingredientId: %s with error %s", ingredientId, err)
	}
	return err
}
