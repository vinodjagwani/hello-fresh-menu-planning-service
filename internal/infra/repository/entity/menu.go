package entity

import (
	"github.com/satori/go.uuid"
	"gorm.io/gorm"
	"time"
)

type Menu struct {
	MenuID           uuid.UUID `gorm:"primary_key"`
	MenuName         string    `gorm:"not null"`
	MenuDescription  string    `gorm:"not null"`
	TotalCockingTime string    `gorm:"not null"`
	Recipe           []Recipe  `gorm:"constraint:OnDelete:CASCADE"`
	Week             int       `gorm:"not null"`
	CreatedAt        time.Time `gorm:"not null"`
	UpdatedAt        time.Time `gorm:"not null"`
}

func (menu *Menu) BeforeCreate(tx *gorm.DB) error {
	tx.Statement.SetColumn("MenuID", uuid.NewV4().String())
	tx.Statement.SetColumn("CreatedAt", time.Now())
	return nil
}

func (menu *Menu) BeforeUpdate(tx *gorm.DB) error {
	if tx.Statement.Changed() {
		tx.Statement.SetColumn("UpdatedAt", time.Now())
	}
	return nil
}
