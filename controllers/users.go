package controllers

import (
	"errors"
	"gin-prac-web/m/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context){
	var input models.ReigisterUser

	if c.Request.Method == "GET"{
		render(c, "register.html", gin.H{})
	}

	if c.Request.Method == "POST"{
		input.Username = c.PostForm("Username")
		input.Password = c.PostForm("Password")
		input.ConfirmPwd = c.PostForm("ConfirmPwd")
		if input.Password != input.ConfirmPwd{
			c.HTML(http.StatusBadRequest, "register.html", gin.H{"error": errors.New("兩次輸入密碼不一致")})
		}
		user := models.User{Username: input.Username, Password: input.Password}
		
		models.DB.Create(&user)
		c.Redirect(http.StatusFound, "/login")//find out how to redirect with data:Use session
		return
	}
}

func Login(c *gin.Context){
	render(c, "login.html", gin.H{})
}