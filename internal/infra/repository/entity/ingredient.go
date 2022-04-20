package entity

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
	"time"
)

type Ingredient struct {
	IngredientID          uuid.UUID `gorm:"primary_key"`
	IngredientName        string    `gorm:"not null"`
	IngredientDescription string    `gorm:"not null"`
	Unit                  int64     `gorm:"not null"`
	RecipeID              string    `gorm:"default:null"`
	CreatedAt             time.Time `gorm:"not null"`
	UpdatedAt             time.Time `gorm:"not null"`
}

func (ingredient *Ingredient) BeforeCreate(tx *gorm.DB) error {
	tx.Statement.SetColumn("IngredientID", uuid.NewV4().String())
	tx.Statement.SetColumn("CreatedAt", time.Now())
	return nil
}

func (ingredient *Ingredient) BeforeUpdate(tx *gorm.DB) error {
	if tx.Statement.Changed() {
		tx.Statement.SetColumn("UpdatedAt", time.Now())
	}
	return nil
}
