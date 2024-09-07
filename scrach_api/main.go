package main

import (
	"fmt"
	"net/http"
	"scrach_api/config"
	"scrach_api/controllers"
	"scrach_api/helper"
	"scrach_api/repository"
	"scrach_api/router"
	"scrach_api/services"
)

func main() {
	fmt.Println("Server Start")

	// database
	db := config.ConnectDatabase()

	// repository
	bookRepository := repository.NewBookRepository(db)

	// service
	bookService := services.NewBookServiceImpl(bookRepository)


	// controller
	bookController := controllers.NewBookController(bookService)




	routes := router.NewRouter(bookController) // create new router
	// create new route with controller

	server := http.Server{Addr: "localhost:8080", Handler: routes} // create new server with registered router
	
	// sering new server with handling error
	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
