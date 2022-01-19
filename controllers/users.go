package controllers

import (
	"errors"
	_"fmt"

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

		session := sessions.DefaultMany(c, "msg")
		
        session.Set("msg", "註冊成功")
        session.Save()
		c.Redirect(http.StatusFound, "/login")//find out how to redirect with data:Use session
		return
	}
}

func Login(c *gin.Context){

	if c.Request.Method == "GET"{
		session1 := sessions.DefaultMany(c, "msg")
		msg := session1.Get("msg")
		session1.Clear()
		session1.Save()
		session2 := sessions.DefaultMany(c, "error")
		err := session2.Get("error")
		session2.Clear()
		session2.Save()
		render(c, "login.html", gin.H{"msg": msg, "error": err})
		
	}
	
	if c.Request.Method == "POST"{
		var user models.User
		var input models.LoginUser

		input.Username = c.PostForm("Username")
		input.Password = c.PostForm("Password")
		if err := models.DB.Where("username = ?", input.Username).First(&user).Error; err != nil{
			c.HTML(http.StatusBadRequest, "login.html", gin.H{"error": errors.New("使用者不存在")})
			return
		}
		if input.Password != user.Password{
			c.HTML(http.StatusBadRequest, "login.html", gin.H{"error": errors.New("密碼錯誤")})
			return
		}
		session := sessions.DefaultMany(c, "isAuth")
		session.Set("isAuth", true)
		session.Save()
		c.Redirect(http.StatusFound, "/books")		
		return
	}
}

func Logout(c *gin.Context){
	session1 := sessions.DefaultMany(c, "isAuth")
	session1.Set("isAuth", false)
	session1.Save()
	session2 := sessions.DefaultMany(c, "msg")
	session2.Clear()
	session2.Set("msg", "登出成功")
	session2.Save()
	c.Redirect(http.StatusFound, "/login")
}