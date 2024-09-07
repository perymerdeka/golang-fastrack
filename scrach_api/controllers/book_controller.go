// controller to manage REST API interface to the business logic
package controllers

import (
	"net/http"
	"scrach_api/data/requests"
	"scrach_api/data/response"
	"scrach_api/helper"
	"scrach_api/services"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type BookController struct {
	BookService services.BookService
}

func NewBookController(bookService services.BookService) *BookController {
	return &BookController{BookService: bookService}
}

func (controller *BookController) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	// implement controller
	bookCreateRequest := requests.BookCreateRequest{}
	helper.ReadRequestBody(request, &bookCreateRequest)

	// return controller to business logic
	controller.BookService.Create(request.Context(), bookCreateRequest)
	webResponse := response.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   nil,
	}
	helper.WriteResponseBody(writer, webResponse)
}

func (controller *BookController) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	bookUpdateRequest := requests.BookUpdateRequest{}

	helper.ReadRequestBody(request, &bookUpdateRequest)
	bookId := params.ByName("bookId")
	id, err := strconv.Atoi(bookId)
	helper.PanicIfError(err)
	bookUpdateRequest.Id = id

	controller.BookService.Update(request.Context(), bookUpdateRequest)
	webResponse := response.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   nil,
	}
	helper.WriteResponseBody(writer, webResponse)
}

func (controller *BookController) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	bookId := params.ByName("bookId")
	id, err := strconv.Atoi(bookId)
	helper.PanicIfError(err)

	controller.BookService.Delete(request.Context(), id)

	// write response status
	writeResponse := response.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   nil,
	}
	helper.WriteResponseBody(writer, writeResponse)
}

func (controller *BookController) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	bookId := params.ByName("bookId")
	id, err := strconv.Atoi(bookId)
	helper.PanicIfError(err)

	result := controller.BookService.FindById(request.Context(), id)

	// write response status
	writeResponse := response.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   result,
	}
	helper.WriteResponseBody(writer, writeResponse)
}

func (controller *BookController) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	result := controller.BookService.FindAll(request.Context())

	// write response status
	writeResponse := response.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   result,
	}
	helper.WriteResponseBody(writer, writeResponse)
}
