package repositories

import (
	"pilatix-api-go/models"

	"github.com/jinzhu/gorm"
)

// CategoryRepository represent repository of category.
type CategoryRepository struct {
	db *gorm.DB
}

// Get return all categories.
func (r *CategoryRepository) Get() []models.Category {
	categories := make([]models.Category, 0)
	r.db.Find(&categories)

	return categories
}

// Find return single category.
func (r *CategoryRepository) Find(id uint) models.Category {
	category := models.Category{}
	r.db.First(&category, id)

	return category
}

// Create creates new category model.
func (r *CategoryRepository) Create(data *models.Category) error {
	return r.db.Create(&data).Error
}

// Update updates category model.
func (r *CategoryRepository) Update(data *models.Category) error {
	return r.db.Save(&data).Error
}

// Delete deletes category by ID.
func (r *CategoryRepository) Delete(id uint) error {
	return r.db.Where("id = ?", id).Delete(models.Category{}).Error
}
