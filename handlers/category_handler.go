package handlers

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/yurichandra/pilatix-api-go/models"
	"github.com/yurichandra/pilatix-api-go/objects"
	"github.com/yurichandra/pilatix-api-go/services"
)

// CategoryHandler represent controller of Category.
type CategoryHandler struct {
	categoryService *services.CategoryService
}

// NewCategoryHandler represent constructing new handler of category.
func NewCategoryHandler(s *services.CategoryService) *CategoryHandler {
	return &CategoryHandler{
		categoryService: s,
	}
}

// GetRoutes return all categories routes available
func (h *CategoryHandler) GetRoutes() chi.Router {
	r := chi.NewRouter()

	r.Get("/", h.Get)
	r.Post("/", h.Create)
	r.Route("/{id}", func(r chi.Router) {
		r.Use(h.Context)
		r.Get("/", h.Find)
		r.Patch("/", h.Update)
		r.Delete("/", h.Delete)
	})

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

// Find handle fetching single category.
func (h *CategoryHandler) Find(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	category, ok := ctx.Value(categoryCtx).(models.Category)
	if !ok {
		render.Render(w, r, sendUnprocessableEntityResponse(""))
		return
	}

	render.Render(w, r, objects.CreateCategoryResponse(category))
}

// Create handle creating new category.
func (h *CategoryHandler) Create(w http.ResponseWriter, r *http.Request) {
	request := objects.CategoryRequest{}
	if err := render.Bind(r, &request); err != nil {
		render.Render(w, r, sendUnprocessableEntityResponse(err.Error()))
		return
	}

	category, err := h.categoryService.Create(request.Name)
	if err != nil {
		render.Render(w, r, sendInternalServerErrorResponse(err.Error()))
		return
	}

	render.Status(r, http.StatusCreated)
	render.Render(w, r, objects.CreateCategoryResponse(category))
}

// Update handle updating a category.
func (h *CategoryHandler) Update(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	category, ok := ctx.Value(categoryCtx).(models.Category)
	if !ok {
		render.Render(w, r, sendUnprocessableEntityResponse(""))
		return
	}

	request := objects.CategoryRequest{}
	if err := render.Bind(r, &request); err != nil {
		render.Render(w, r, sendUnprocessableEntityResponse(err.Error()))
		return
	}

	newCategory, err := h.categoryService.Update(category.ID, request.Name)
	if err != nil {
		render.Render(w, r, sendInternalServerErrorResponse(err.Error()))
		return
	}

	render.Render(w, r, objects.CreateCategoryResponse(newCategory))
}

// Delete handle deleting a category
func (h *CategoryHandler) Delete(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	category, ok := ctx.Value(categoryCtx).(models.Category)
	if !ok {
		render.Render(w, r, sendUnprocessableEntityResponse(""))
		return
	}

	err := h.categoryService.Delete(category.ID)
	if err != nil {
		render.Render(w, r, sendInternalServerErrorResponse(err.Error()))
		return
	}

	render.JSON(w, r, "")
}
