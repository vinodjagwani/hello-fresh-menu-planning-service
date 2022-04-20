package entity

import (
	"github.com/satori/go.uuid"
	"gorm.io/gorm"
	"time"
)

type Recipe struct {
	RecipeID             uuid.UUID `gorm:"primary_key"`
	RecipeName           string    `gorm:"not null"`
	RecipeDescription    string
	RecipeInstructions   string
	MenuID               string `gorm:"default:null"`
	RecipeClassification string
	Ingredient           []Ingredient `gorm:"constraint:OnDelete:CASCADE"`
	CreatedAt            time.Time    `gorm:"not null"`
	UpdatedAt            time.Time    `gorm:"not null"`
}

func (recipe *Recipe) BeforeCreate(tx *gorm.DB) error {
	tx.Statement.SetColumn("RecipeID", uuid.NewV4().String())
	tx.Statement.SetColumn("CreatedAt", time.Now())
	return nil
}

func (recipe *Recipe) BeforeUpdate(tx *gorm.DB) error {
	if tx.Statement.Changed() {
		tx.Statement.SetColumn("UpdatedAt", time.Now())
	}
	return nil
}
