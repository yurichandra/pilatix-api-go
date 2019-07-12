package services

import (
	"errors"
	"pilatix-api-go/models"
	"pilatix-api-go/repositories"
)

// CategoryService represent service layer of Category.
type CategoryService struct {
	categoryRepo repositories.CategoryRepositoryContract
}

// Get return all categories.
func (s *CategoryService) Get() []models.Category {
	return s.categoryRepo.Get()
}

// Create creates new Category data.
func (s *CategoryService) Create(name string) (models.Category, error) {
	category := models.Category{
		Name: name,
	}

	if err := s.categoryRepo.Create(&category); err != nil {
		return models.Category{}, err
	}

	return category, nil
}

// Find finds Category data.
func (s *CategoryService) Find(id uint) (models.Category, error) {
	category := s.categoryRepo.Find(id)
	if category.ID == 0 {
		return models.Category{}, errors.New("Fail to find category with certain ID")
	}

	return category, nil
}

// Update updates Category data.
func (s *CategoryService) Update(id uint, name string) (models.Category, error) {
	updateCategory := models.Category{
		ID:   id,
		Name: name,
	}

	if err := s.categoryRepo.Update(&updateCategory); err != nil {
		return models.Category{}, err
	}

	return updateCategory, nil
}

// Delete deletes Category data.
func (s *CategoryService) Delete(id uint) error {
	return s.categoryRepo.Delete(id)
}
