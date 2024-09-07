package router

import (
	"fmt"
	"net/http"
	"scrach_api/controllers"

	"github.com/julienschmidt/httprouter"
)

func NewRouter(bookController *controllers.BookController) *httprouter.Router {
	router := httprouter.New()

	router.GET("/", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		fmt.Fprint(w, "Hello World")
	})

	router.GET("/api/books", bookController.FindAll)
	router.GET("/api/books/:bookId", bookController.FindById)

	router.POST("/api/books", bookController.Create)
	router.PATCH("/api/books", bookController.Update)
	router.DELETE("/api/books/:bookId", bookController.Delete)

	return router


}