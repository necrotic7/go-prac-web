package main

import (
	_"gin-prac-web/m/controllers"
	"gin-prac-web/m/models"
	_ "net/http"

	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func main(){
	Router = gin.Default()
	models.ConnectDataBase()
	Router.LoadHTMLGlob("template/*")
	initializeRoutes()
	Router.Run()
}