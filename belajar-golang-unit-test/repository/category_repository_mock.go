package repository

import (
	"belajar-golang-unit-test/entity"

	"github.com/stretchr/testify/mock"
)

type CategoryRepositoryMock struct {
	// Mock ini adalah field yang berisi object Mock/struct Mock
	Mock mock.Mock
}

// Disinilah kita mencoba mengimplementasikan function mengambil data yang disebut FindById
// Ini adalah function untuk pura pura/untuk simulasi pengambilan data ke database yang
// nantinya akan dieksekusi oleh mock

// Saya akan menyimpan function findById di pointernya CategoryRepositoryMock
func (repository *CategoryRepositoryMock) FindById(id string) *entity.Category {
	// Disini ketika Mock dieksekusi mock akan melakukan pencarian ke call stack
	// dan tahu bahwa yang memanggilnya adalah function FindById, nah karena sebelumnya sudah
	// didaftarkan Skenarionya lewat Mock.On maka Mock.Called akan mencocokan ke daftarnya
	// Jika cocok maka akan ditandai sebagai terpenuhi dan mengembalikan nilai sesuai yang didaftarkan
	// tapi jika tidak Mock akan Panic
	arguments := repository.Mock.Called(id)
	if arguments.Get(0) == nil {
		return nil
	} else {
		// karena yang dikemabilkan berpa interface{} kosong maka bisa gunakan type assert
		category := arguments.Get(0).(entity.Category)
		// Karena tipe yang harus dikembalikan oleh function ini berupa pointer entity.Category
		// maka perlu ditambahkan & dan pastikan variablenya bertipe entity.Category
		return &category
	}
}
