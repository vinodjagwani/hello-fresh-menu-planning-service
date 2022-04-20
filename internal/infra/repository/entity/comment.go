package entity

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
	"time"
)

type Comment struct {
	CommentID  uuid.UUID `gorm:"primary_key"`
	Comment    string    `gorm:"not null"`
	MenuPlanID string    `gorm:"not null"`
	UserID     string    `gorm:"not null"`
	CreatedAt  time.Time `gorm:"not null"`
	UpdatedAt  time.Time `gorm:"not null"`
}

func (comment *Comment) BeforeCreate(tx *gorm.DB) error {
	tx.Statement.SetColumn("CommentID", uuid.NewV4().String())
	tx.Statement.SetColumn("CreatedAt", time.Now())
	return nil
}

func (comment *Comment) BeforeUpdate(tx *gorm.DB) error {
	if tx.Statement.Changed() {
		tx.Statement.SetColumn("UpdatedAt", time.Now())
	}
	return nil
}
