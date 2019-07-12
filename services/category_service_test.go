package services

import (
	"pilatix-api-go/mocks"
	"pilatix-api-go/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetCategory(t *testing.T) {
	mockCategoryRepo := &mocks.CategoryRepositoryMock{}
	mockCategoryService := &CategoryService{
		categoryRepo: mockCategoryRepo,
	}

	expected := []models.Category{
		models.Category{
			ID:   1,
			Name: _testFaker.Lorem().Word(),
		},
		models.Category{
			ID:   2,
			Name: _testFaker.Lorem().Word(),
		},
		models.Category{
			ID:   3,
			Name: _testFaker.Lorem().Word(),
		},
	}

	mockCategoryRepo.On("Get").Return(expected)
	actual := mockCategoryService.Get()

	assert.Equal(t, expected, actual)
}

func TestCreateCategory(t *testing.T) {
	mockCategoryRepo := &mocks.CategoryRepositoryMock{}
	mockCategoryService := &CategoryService{
		categoryRepo: mockCategoryRepo,
	}

	category := models.Category{
		Name: _testFaker.Lorem().Word(),
	}

	mockCategoryRepo.On("Create", &category).Return(nil)
	expected, err := mockCategoryService.Create(category.Name)

	assert.Nil(t, err)
	assert.Equal(t, expected.Name, category.Name)
}

func TestFindCategory(t *testing.T) {
	mockCategoryRepo := &mocks.CategoryRepositoryMock{}
	mockCategoryService := &CategoryService{
		categoryRepo: mockCategoryRepo,
	}

	expected := models.Category{
		ID:   1,
		Name: _testFaker.Lorem().Word(),
	}

	mockCategoryRepo.On("Find", expected.ID).Return(expected)
	actual, err := mockCategoryService.Find(expected.ID)

	assert.Nil(t, err)
	assert.Equal(t, expected, actual)
	assert.Equal(t, expected.ID, actual.ID)
	assert.Equal(t, expected.Name, actual.Name)
}

func TestUpdateCategory(t *testing.T) {
	mockCategoryRepo := &mocks.CategoryRepositoryMock{}
	mockCategoryService := &CategoryService{
		categoryRepo: mockCategoryRepo,
	}

	expected := models.Category{
		ID:   1,
		Name: _testFaker.Lorem().Word(),
	}

	mockCategoryRepo.On("Update", &expected).Return(nil)
	actual, err := mockCategoryService.Update(expected.ID, expected.Name)

	assert.Nil(t, err)
	assert.Equal(t, expected, actual)
	assert.Equal(t, expected.ID, actual.ID)
	assert.Equal(t, expected.Name, actual.Name)
}

func TestDeleteCategory(t *testing.T) {
	mockCategoryRepo := &mocks.CategoryRepositoryMock{}
	mockCategoryService := &CategoryService{
		categoryRepo: mockCategoryRepo,
	}

	category := models.Category{
		ID:   1,
		Name: _testFaker.Lorem().Word(),
	}

	mockCategoryRepo.On("Delete", category.ID).Return(nil)
	err := mockCategoryService.Delete(category.ID)

	assert.Nil(t, err)
}
