package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"pilatix-api-go/db"
	"pilatix-api-go/handlers"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

func initRouter() chi.Router {
	r := chi.NewRouter()

	r.Use(render.SetContentType(render.ContentTypeJSON))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		payload := map[string]interface{}{
			"name":    "pilatix-api",
			"version": "1",
		}

		res, _ := json.Marshal(payload)
		w.Write(res)
	})

	r.Mount("/categories", handlers.NewCategoryHandler(categoryService).GetRoutes())

	return r
}

func serveHTTP() {
	router := initRouter()
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8000"
	}

	fmt.Printf("App running on port %s\n", port)
	http.ListenAndServe(fmt.Sprintf(":%s", port), router)
}

func runMigration() {
	fmt.Println("Running migration")
	db.Migrate()
	fmt.Println("Migration completed!")
}
