package handlers

import (
	"context"
	"fmt"
	"net/http"
	"pilatix-api-go/objects"
	"pilatix-api-go/services"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

// CategoryHandler represent controller of Category.
type CategoryHandler struct {
	categoryService services.CategoryService
}

// NewCategoryHandler represent constructing new handler of category.
func NewCategoryHandler(s *services.CategoryService) *CategoryHandler {
	return &CategoryHandler{
		categoryService: *s,
	}
}

// GetRoutes return all categories routes available
func (h *CategoryHandler) GetRoutes() chi.Router {
	r := chi.NewRouter()

	r.Get("/", h.Get)

	return r
}

// Context is defining context of route params.
func (h *CategoryHandler) Context(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id, _ := strconv.Atoi(chi.URLParam(r, "id"))
		category, err := h.categoryService.Find(uint(id))
		if err != nil {
			render.Render(w, r, sendNotFoundResponse(err.Error()))
			return
		}

		ctx := context.WithValue(r.Context(), categoryCtx, category)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// Get handle fetching all categories.
func (h *CategoryHandler) Get(w http.ResponseWriter, r *http.Request) {
	categoryListResponse := objects.CreateCategoryListResponse(h.categoryService.Get())

	if err := render.RenderList(w, r, categoryListResponse); err != nil {
		fmt.Println(err.Error())
		return
	}
}
