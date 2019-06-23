package objects

import (
	"errors"
	"net/http"
)

// CategoryRequest represent object request of Category model
type CategoryRequest struct {
	Name string `json:"name"`
}

// Validate doing validation of request object
func (req *CategoryRequest) Validate(r *http.Request) error {
	if req.Name == "" {
		return errors.New("Field `name` can't be empty")
	}

	return nil
}
