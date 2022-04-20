package entity

import (
	"gorm.io/gorm"
)

func PaginateScope(p int, ps int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		page := p
		if page == 0 {
			page = 1
		}
		pageSize := ps
		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}
		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}
