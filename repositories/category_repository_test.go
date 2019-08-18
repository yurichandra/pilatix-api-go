package repositories

import (
	"testing"

	"github.com/yurichandra/pilatix-api-go/db"
	"github.com/yurichandra/pilatix-api-go/models"
)

func TestGetCategory(t *testing.T) {
	db.Reset()
	for i := 0; i < 5; i++ {
		_testDB.Create(&models.Category{
			Name: _testFaker.Lorem().Word(),
		})
	}

	testCategoryRepo := &CategoryRepository{
		db: _testDB,
	}

	expected := testCategoryRepo.Get()
	if len(expected) != 5 {
		t.Errorf("Length of categories not equal to 5")
	}
}

func TestFindCategory(t *testing.T) {
	db.Reset()
	testCategoryRepo := &CategoryRepository{
		db: _testDB,
	}

	mockedCategory := models.Category{
		Name: _testFaker.Lorem().Word(),
	}

	err := testCategoryRepo.Create(&mockedCategory)
	if err != nil {
		t.Errorf("Error occured when run test find category")
	}

	expected := testCategoryRepo.Find(mockedCategory.ID)
	if expected.Name != mockedCategory.Name {
		t.Errorf("Find category fail to find a category by ID")
	}
}

func TestCreateCategory(t *testing.T) {
	db.Reset()
	testCategoryRepo := &CategoryRepository{
		db: _testDB,
	}

	mockedCategory := models.Category{
		Name: _testFaker.Lorem().Word(),
	}

	err := testCategoryRepo.Create(&mockedCategory)
	if err != nil {
		t.Errorf("Test fail on create new data")
	}
}

func TestUpdateCategory(t *testing.T) {
	db.Reset()
	testCategoryRepo := &CategoryRepository{
		db: _testDB,
	}

	mockedCategory := models.Category{
		Name: _testFaker.Lorem().Word(),
	}

	err := testCategoryRepo.Create(&mockedCategory)
	if err != nil {
		t.Errorf("Fail to create on update process")
	}

	updateMockedCategory := models.Category{
		Name: _testFaker.Lorem().Word(),
	}

	newErr := testCategoryRepo.Update(&updateMockedCategory)
	if newErr != nil {
		t.Errorf("Fail to update category data")
	}
}

func TestDeleteCategory(t *testing.T) {
	db.Reset()
	testCategoryRepo := CategoryRepository{
		db: _testDB,
	}

	mockedCategory := models.Category{
		Name: _testFaker.Lorem().Word(),
	}

	testCategoryRepo.Create(&mockedCategory)

	err := testCategoryRepo.Delete(mockedCategory.ID)
	if err != nil {
		t.Errorf("Fail to delete category data")
	}
}
