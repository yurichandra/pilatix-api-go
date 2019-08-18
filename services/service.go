package services

import (
	"github.com/yurichandra/pilatix-api-go/repositories"
)

// NewCategoryService represent initialization of CategoryService.
func NewCategoryService(repo repositories.CategoryRepositoryContract) *CategoryService {
	return &CategoryService{
		categoryRepo: repo,
	}
}
