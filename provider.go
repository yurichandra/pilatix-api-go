package main

import (
	"log"
	"pilatix-api-go/db"
	"pilatix-api-go/repositories"
	"pilatix-api-go/services"
)

var categoryService *services.CategoryService

func boot() {
	var err error

	dbConn := db.Get()

	categoryService = services.NewCategoryService(repositories.NewCategoryRepository(dbConn))

	if err != nil {
		log.Fatal(err)
	}
}
