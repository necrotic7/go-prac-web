package main

import (
	_ "gin-prac-web/m/controllers"
	"gin-prac-web/m/models"
	_ "net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

var Router *gin.Engine


func main(){
	Router = gin.Default()
	models.ConnectDataBase()
	Router.LoadHTMLGlob("template/*")
	store := cookie.NewStore([]byte("msg"))
	Router.Use(sessions.Sessions("my", store))
	
	initializeRoutes()
	Router.Run()
}

