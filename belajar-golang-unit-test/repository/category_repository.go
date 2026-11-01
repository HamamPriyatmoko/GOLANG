package repository

import "belajar-golang-unit-test/entity"

type CategoryRepository interface {
	FindById(Id string) *entity.Category
}
