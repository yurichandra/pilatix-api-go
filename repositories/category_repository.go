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
func (repo *CategoryRepository) Get() []models.Category {
	categories := make([]models.Category, 0)
	repo.db.Find(&categories)

	return categories
}

// Find return single category.
func (repo *CategoryRepository) Find(id uint) models.Category {
	category := models.Category{}
	repo.db.First(&category, id)

	return category
}

// Create creates new category model.
func (repo *CategoryRepository) Create(data *models.Category) error {
	return repo.db.Create(&data).Error
}

// Update updates category model.
func (repo *CategoryRepository) Update(data *models.Category) error {
	return repo.db.Save(&data).Error
}

// Delete deletes category by ID.
func (repo *CategoryRepository) Delete(id uint) error {
	return repo.db.Where("id = ?", id).Delete(models.Category{}).Error
}
