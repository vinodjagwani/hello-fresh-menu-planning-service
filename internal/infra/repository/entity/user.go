package entity

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
	"time"
)

type User struct {
	UserID    uuid.UUID `gorm:"primaryKey"`
	UserName  string    `gorm:"not null"`
	Email     string    `gorm:"not null; unique"`
	UserType  string    `gorm:"not null"`
	Password  string    `gorm:"not null"`
	CreatedAt time.Time `gorm:"not null"`
	UpdatedAt time.Time `gorm:"not null"`
}

func (user *User) BeforeCreate(tx *gorm.DB) error {
	tx.Statement.SetColumn("UserID", uuid.NewV4().String())
	tx.Statement.SetColumn("CreatedAt", time.Now())
	return nil
}

func (user *User) BeforeUpdate(tx *gorm.DB) error {
	if tx.Statement.Changed() {
		tx.Statement.SetColumn("UpdatedAt", time.Now())
	}
	return nil
}
