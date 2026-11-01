package service

import (
	"belajar-golang-unit-test/entity"
	"belajar-golang-unit-test/repository"
	"errors"
)

type CategoryService struct {
	Repository repository.CategoryRepository
}

// Ini adalah function yang ingin kita lakukan test
func (service CategoryService) Get(id string) (*entity.Category, error) {
	// FindById diibaratkan yang akan ambil data from database(Entity)
	category := service.Repository.FindById(id)
	if category == nil {
		return nil, errors.New("Category Not Found")
	} else {
		return category, nil
	}
}
