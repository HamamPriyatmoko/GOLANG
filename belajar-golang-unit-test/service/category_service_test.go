package service

import (
	"belajar-golang-unit-test/entity"
	"belajar-golang-unit-test/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Kita akan membuat object pointer dari CategoryRepositoryMock dan ini sudah mengikuti kontrak dari
// interface CategoryRepository
var categoryRepository = &repository.CategoryRepositoryMock{Mock: mock.Mock{}}

// Nah disini kita membuat object CategoryService yang berisi Repository
// dimana Repository ini bertipe interface CategoryRepository nah kenapa Repository
// diisi dengan variable categoryRepository adalah karena categoryRepository bertipe CategoryRepositoryMock
// dimana sudah memenuhi kontrak interface CategoryRepository
var categoryService = CategoryService{Repository: categoryRepository}

// testing.T adalah sebuah tipe untuk membuat unit test
func TestCategoryService_GetNotFound(t *testing.T) {
	// Mock On itu ibarat listener dijavascript jadi ketika ada Function tertentu yang dipanggil
	// maka Mock akan melakukan eksekusi berikut, lalu kita bisa set kembalian dari findById jika
	// dipanggil menggunakan method Return
	categoryRepository.Mock.On("FindById", "1").Return(nil)
	// categoryRepository.Mock.On("FindById", "2").Return(category)

	result, err := categoryService.Get("1")

	assert.Nil(t, result)
	assert.NotNil(t, err)
}

func TestCategoryService_GetSucces(t *testing.T) {
	category := entity.Category{
		Id:   "2",
		Name: "Hamam",
	}

	// Mock On itu ibarat listener dijavascript jadi ketika ada Function tertentu yang dipanggil
	// maka Mock akan melakukan eksekusi berikut, lalu kita bisa set kembalian dari findById jika
	// dipanggil menggunakan method Return
	// categoryRepository.Mock.On("FindById", "1").Return(nil)
	categoryRepository.Mock.On("FindById", "2").Return(category)

	result, err := categoryService.Get("2")

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, category.Id, result.Id)
	assert.Equal(t, category.Name, result.Name)
}
