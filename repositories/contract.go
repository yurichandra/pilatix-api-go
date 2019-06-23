package repositories

import (
	"pilatix-api-go/models"
)

// CategoryRepositoryContract represent contract of CategoryRepository.
type CategoryRepositoryContract interface {
	Get() []models.Category
	Find(id uint) models.Category
	Create(data *models.Category) error
	Update(data *models.Category) error
	Delete(id uint) error
}
