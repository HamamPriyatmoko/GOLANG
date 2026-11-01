package helper

import (
	"HamamPriyatmoko/belajar-golang-restful-api/model/domain"
	"HamamPriyatmoko/belajar-golang-restful-api/model/web"
)

func ToCategoryResponse(category domain.Category) web.CategoryResponse {
	return web.CategoryResponse{
		Id:   category.Id,
		Name: category.Name,
	}
}
