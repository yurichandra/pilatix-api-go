package repositories

import (
	"github.com/jinzhu/gorm"
)

// NewCategoryRepository creates new CategoryRepository.
func NewCategoryRepository(db *gorm.DB) *CategoryRepository {
	return &CategoryRepository{
		db: db,
	}
}
