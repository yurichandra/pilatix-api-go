package services

import (
	"pilatix-api-go/models"
)

// CategoryServiceContract represent contract of CategoryService
type CategoryServiceContract interface {
	Get() []models.Category
	Create(name string) (models.Category, error)
	Find(id uint) (models.Category, error)
	Update(id uint, name string) (models.Category, error)
	Delete(id uint) error
}
