package handlers

import (
	"net/http"

	"github.com/go-chi/render"
)

type errorResponse struct {
	HTTPStatus int    `json:"-"`
	Message    string `json:"error_message"`
}

func (e *errorResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HTTPStatus)
	return nil
}

func sendNotFoundResponse(message string) *errorResponse {
	if message == "" {
		message = "Resource not found"
	}

	return &errorResponse{
		HTTPStatus: http.StatusNotFound,
		Message:    message,
	}
}
