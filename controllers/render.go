package controllers

import (
	
	"net/http"

	"github.com/gin-gonic/gin"
)

func render(c *gin.Context,templateName string, data gin.H ){
	
	switch c.Request.Header.Get("Accept"){
	case "application/json":
		c.JSON(http.StatusOK, data["data"])
	case "application/xml":
		c.XML(http.StatusOK, data["data"])
	default:
		c.HTML(http.StatusOK, templateName, data)
	}
}