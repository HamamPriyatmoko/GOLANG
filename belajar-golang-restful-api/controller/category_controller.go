package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// interface bersifat seperti kontrak jadi semua harus terpenuhi contohnya jika kita menuliskan didalam kontraknya
// ada method Create,Update,Delete maka variable/struct yang mau mengimplementasikan interface ini harus mempunyai
// method tersebut
type CategoryController interface {
	Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}
