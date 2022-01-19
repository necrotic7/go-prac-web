package controllers

import (
	"fmt"
	"gin-prac-web/m/models"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func BooksIndex(c *gin.Context){
	session := sessions.DefaultMany(c, "isAuth")
	Auth := session.Get("isAuth")
	fmt.Println(Auth)
	if Auth.(bool) {
		var books []models.Book
		models.DB.Find(&books)
		render(c, "index.html", gin.H{"data": books})
	}else{
		session2 := sessions.DefaultMany(c, "error")
		session2.Clear()
		session2.Set("error", "請先登入")
		session2.Save()
		c.Redirect(http.StatusFound, "/login")
	}
	
}


func AddBook(c *gin.Context){
	if method := c.Request.Method; method == "GET"{
		render(c, "add.html", gin.H{})
		return
	}
	
	var input models.AddBookInput

	if form := c.PostForm("Title"); form != ""{
		input.Title = c.PostForm("Title")
		input.Author = c.PostForm("Author")
		book := models.Book{Title: input.Title, Author: input.Author}
		fmt.Println(input)
		models.DB.Create(&book)
		c.Redirect(http.StatusFound, "/books")//find what status can pass into redirect
		return
	}else{
		// ShouldBindJSON:  If the data is invalid, it will return a 400 error 
		// to the client and tell them which fields are invalid.
		if err:= c.ShouldBindJSON(&input); err != nil{
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	
		book := models.Book{Title: input.Title, Author: input.Author}
		models.DB.Create(&book)
	
		render(c, "index.html", gin.H{"data": book,})
	}
	
}

func FindBook(c *gin.Context){
	var book models.Book

	if err := models.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error":"not found"})
		return
	}

	render(c, "article.html", gin.H{"data":book})
}

func UpdateBook(c *gin.Context) {
	// Get model if exist
	var book models.Book
	var input models.UpdateBookInput
	if err := models.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	
	//show form page
	if method := c.Request.Method; method == "GET"{
		render(c, "update.html", gin.H{"data":book})
		return
	}
	//pass form data
	if method := c.Request.Method; method == "POST"{
		input.Title = c.PostForm("Title")
		input.Author = c.PostForm("Author")
		
		models.DB.Model(&book).Updates(models.Book{Title:input.Title, Author:input.Author})
		c.Redirect(http.StatusFound, "/books")//find what status can pass into redirect
		return
	//for API PATCH
	}else if method := c.Request.Method; method == "PATCH"{
		
		// Validate input
		var input models.UpdateBookInput
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		
		uperr:=models.DB.Model(&book).Updates(models.Book{Title:input.Title, Author:input.Author}).Error
		if uperr != nil{
			panic(uperr)
		}
		c.JSON(http.StatusOK, gin.H{"data": book})
	}
}

func DeleteBook(c *gin.Context){
	var book models.Book
	if err := models.DB.Where("id = ? ", c.Param("id")).First(&book).Error; err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": "not found"})
	}

	models.DB.Delete(&book)
	c.Redirect(http.StatusFound, "/books")
}