package db

import (
	"fmt"
	"os"

	"github.com/yurichandra/pilatix-api-go/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" //Postgres dialect.
)

var db *gorm.DB

func initDB() (*gorm.DB, error) {
	connectionString := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_DATABASE"),
		os.Getenv("DB_PASSWORD"),
	)

	return gorm.Open("postgres", connectionString)
}

// Get return database instance.
func Get() *gorm.DB {
	if db == nil {
		conn, err := initDB()
		if err != nil {
			panic(err.Error())
		}

		db = conn
	}

	return db
}

// Migrate will migrating model to database.
func Migrate() {
	Get().AutoMigrate(
		&models.Category{},
	)
}

// Drop removes exist table
func Drop() {
	Get().DropTableIfExists(
		&models.Category{},
	)
}

// Reset do migration and drop table
func Reset() {
	Drop()
	Migrate()
}
