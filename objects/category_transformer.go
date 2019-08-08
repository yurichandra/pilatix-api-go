package objects

import (
	"pilatix-api-go/models"

	"github.com/go-chi/render"
)

// CreateCategoryResponse is transformer of Category model.
func CreateCategoryResponse(category models.Category) render.Renderer {
	return &CategoryResponse{
		ID:        category.ID,
		Name:      category.Name,
		CreatedAt: category.CreatedAt.Format("2006-02-10 15:04:05"),
		UpdatedAt: category.UpdatedAt.Format("2006-02-10 15:04:05"),
	}
}

// CreateCategoryListResponse is transformer of list categories
func CreateCategoryListResponse(data []models.Category) []render.Renderer {
	categories := make([]render.Renderer, 0)

	for _, category := range data {
		categories = append(categories, CreateCategoryResponse(category))
	}

	return categories
}
