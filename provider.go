package main

import (
	"log"

	"github.com/yurichandra/pilatix-api-go/db"
	"github.com/yurichandra/pilatix-api-go/repositories"
	"github.com/yurichandra/pilatix-api-go/services"
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
