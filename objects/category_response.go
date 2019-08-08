package objects

import "net/http"

// CategoryResponse represent response of Category object
type CategoryResponse struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

// Render is rendering object of Category.
func (c *CategoryResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
