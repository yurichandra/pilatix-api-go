package repositories

import (
	"fmt"
	"testing"

	"github.com/jinzhu/gorm"
	"github.com/yurichandra/pilatix-api-go/db"

	"github.com/jaswdr/faker"
	"github.com/joho/godotenv"
)

var (
	_testFaker faker.Faker
	_testDB    *gorm.DB
)

func TestMain(t *testing.M) {
	err := godotenv.Load("../.env.test")
	if err != nil {
		fmt.Println(".env.test was not specified")
	}

	_testFaker = faker.New()
	_testDB = db.Get()

	t.Run()
}
