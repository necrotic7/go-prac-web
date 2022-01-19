package controllers

import (
	"errors"
	"gin-prac-web/m/models"
	
	"net/http"
	"github.com/gin-contrib/sessions"

	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
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
			return
		}
		
		user := models.User{Username: input.Username, Password: input.Password}
		
		var mysqlerror *mysql.MySQLError
		err := models.DB.Create(&user)
		if err != nil{
			if errors.As(err.Error, &mysqlerror) && mysqlerror.Number == 1062{
				c.HTML(http.StatusBadRequest, "register.html", gin.H{"error": errors.New("使用者已存在")})
				return
			}
		}

		session := sessions.Default(c)
		
        session.Set("msg", "註冊成功")
        session.Save()
		c.Redirect(http.StatusFound, "/login")//find out how to redirect with data:Use session
		return
	}
}

func Login(c *gin.Context){
	session := sessions.Default(c)
    msg := session.Get("msg")
	session.Clear()
	session.Save()
	render(c, "login.html", gin.H{"msg": msg})
	
}