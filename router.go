package main

import (
	"gin-prac-web/m/controllers"
)

func initializeRoutes() {

	Router.GET("/books", controllers.BooksIndex)

	Router.GET("/books/add", controllers.AddBook)
	Router.POST("/books/add", controllers.AddBook)

	Router.GET("/books/:id", controllers.FindBook)

	Router.GET("/books/edit/:id", controllers.UpdateBook)
	Router.POST("/books/edit/:id", controllers.UpdateBook)
	Router.PATCH("/books/edit/:id", controllers.UpdateBook)
	
	Router.GET("/books/delete/:id", controllers.DeleteBook)
	Router.DELETE("/books/delete/:id", controllers.DeleteBook)
}
