package controller

import (
	"HamamPriyatmoko/belajar-golang-restful-api/helper"
	"HamamPriyatmoko/belajar-golang-restful-api/model/web"
	"HamamPriyatmoko/belajar-golang-restful-api/service"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type CategoryControllerImpl struct {
	CategoryService service.CategoryService
}

// karena CategoryControllerImpl sudah mengimplementasikan method sesuai dengan contract interface
// CategoryControllernya maka CategoryControllerImpl termasuk dalam type interface CategoryController
func NewCategoryController(categoryService service.CategoryService) CategoryController {
	return &CategoryControllerImpl{
		CategoryService: categoryService,
	}
}

func (controller *CategoryControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryCreateRequest := web.CategoryCreateRequest{}
	helper.ReadFromRequestBody(request, &categoryCreateRequest)

	// Kenapa pake context dari request, alih-alih menggunakan context yang kita buat sendiri context request
	// lebih baik karena jika user membatalkan koneksi/tutup web browsernya maka request akan langsung memberikan
	// sinyal batal menggunakan context sehingga service tidak harus terus berjalan jika user sudah tidak
	// ingin menjalankannya lagi
	categoryResponse := controller.CategoryService.Create(request.Context(), categoryCreateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *CategoryControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryUpdateRequest := web.CategoryUpdateRequest{}
	helper.ReadFromRequestBody(request, &categoryUpdateRequest)

	categoryId := params.ByName("categoryId")
	id, err := strconv.Atoi(categoryId)
	helper.PanicIfError(err)

	categoryUpdateRequest.Id = id

	// Kenapa pake context dari request, alih-alih menggunakan context yang kita buat sendiri context request
	// lebih baik karena jika user membatalkan koneksi/tutup web browsernya maka request akan langsung memberikan
	// sinyal batal menggunakan context sehingga service tidak harus terus berjalan jika user sudah tidak
	// ingin menjalankannya lagi
	categoryResponse := controller.CategoryService.Update(request.Context(), categoryUpdateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *CategoryControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	// method ByName dari httprouter.Params hanya mengembalikan string, jadi perlu adanya conversi
	categoryId := params.ByName("categoryId")
	id, err := strconv.Atoi(categoryId)
	helper.PanicIfError(err)

	// Kenapa pake context dari request, alih-alih menggunakan context yang kita buat sendiri context request
	// lebih baik karena jika user membatalkan koneksi/tutup web browsernya maka request akan langsung memberikan
	// sinyal batal menggunakan context sehingga service tidak harus terus berjalan jika user sudah tidak
	// ingin menjalankannya lagi
	controller.CategoryService.Delete(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *CategoryControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	// method ByName dari httprouter.Params hanya mengembalikan string, jadi perlu adanya conversi
	categoryId := params.ByName("categoryId")
	id, err := strconv.Atoi(categoryId)
	helper.PanicIfError(err)

	// Kenapa pake context dari request, alih-alih menggunakan context yang kita buat sendiri context request
	// lebih baik karena jika user membatalkan koneksi/tutup web browsernya maka request akan langsung memberikan
	// sinyal batal menggunakan context sehingga service tidak harus terus berjalan jika user sudah tidak
	// ingin menjalankannya lagi
	categoryResponse := controller.CategoryService.FindById(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *CategoryControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoriesResponse := controller.CategoryService.FindAll(request.Context())

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoriesResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
