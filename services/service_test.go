package services

import (
	"pilatix-api-go/repositories"
	"testing"

	"github.com/jaswdr/faker"
)

var (
	_testFaker        faker.Faker
	_testCategoryRepo *repositories.CategoryRepository
)

func TestMain(t *testing.M) {
	_testFaker = faker.New()

	t.Run()
}
