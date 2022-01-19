package main

import (
	"gin-prac-web/m/controllers"
	_"github.com/gin-contrib/sessions"
	_"github.com/gin-contrib/sessions/cookie"
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

	Router.GET("/register", controllers.Register)
	Router.POST("/register", controllers.Register)
	

	Router.GET("/login", controllers.Login)
	Router.POST("/login", controllers.Login)
	Router.GET("/logout", controllers.Logout)
}
