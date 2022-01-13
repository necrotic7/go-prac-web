package main

import (
	"gin-prac-web/m/controllers"
	"gin-prac-web/m/models"
	_ "net/http"

	"github.com/gin-gonic/gin"
)

func main(){
	router := gin.Default()
	models.ConnectDataBase()

	router.GET("/books", controllers.FindBooks)
	router.POST("/books", controllers.AddBook)
	router.GET("/books/:id", controllers.FindBook)
	router.PATCH("/books/:id", controllers.UpdateBook)
	router.DELETE("/books/:id", controllers.DeleteBook)

	router.Run()
}