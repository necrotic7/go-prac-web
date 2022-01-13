package main

import "gin-prac-web/m/controllers"

func initializeRoutes() {

	Router.GET("/books", controllers.FindBooks)
	Router.POST("/books", controllers.AddBook)
	Router.GET("/books/:id", controllers.FindBook)
	Router.PATCH("/books/:id", controllers.UpdateBook)
	Router.DELETE("/books/:id", controllers.DeleteBook)
}
