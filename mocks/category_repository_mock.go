package mocks

import (
	"github.com/yurichandra/pilatix-api-go/models"

	"github.com/stretchr/testify/mock"
)

// CategoryRepositoryMock represent mocking of CategoryRepository.
type CategoryRepositoryMock struct {
	mock.Mock
}

// Get mock Get method of CategoryRepository.
func (r *CategoryRepositoryMock) Get() []models.Category {
	args := r.Called()

	return args.Get(0).([]models.Category)
}

// Find mock Find method of CategoryRepository.
func (r *CategoryRepositoryMock) Find(id uint) models.Category {
	args := r.Called(id)

	return args.Get(0).(models.Category)
}

// Create mock Create method of CategoryRepository.
func (r *CategoryRepositoryMock) Create(data *models.Category) error {
	args := r.Called(data)

	return args.Error(0)
}

// Update mock Update method of CategoryRepository.
func (r *CategoryRepositoryMock) Update(data *models.Category) error {
	args := r.Called(data)

	return args.Error(0)
}

// Delete mock Delete method of CategoryRepository.
func (r *CategoryRepositoryMock) Delete(id uint) error {
	args := r.Called(id)

	return args.Error(0)
}
