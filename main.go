package main

import (
	"gin-bookstore/m/controllers"
	"gin-bookstore/m/models"
	_ "net/http"

	"github.com/gin-gonic/gin"
)

func main(){
	router := gin.Default()
	models.ConnectDataBase()

	router.GET("/books", controllers.FindBooks)
	router.POST("/books", controllers.AddBook)
	router.GET("/books/:id", controllers.FindBook)
	
	router.Run()
}