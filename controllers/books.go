package controllers

import (
	"fmt"
	"gin-prac-web/m/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func BooksIndex(c *gin.Context){
	var books []models.Book
	models.DB.Find(&books)

	c.HTML(http.StatusOK, "index.html", gin.H{"data": books})
}

func AddBook(c *gin.Context){
	var input models.AddBookInput
	// ShouldBindJSON:  If the data is invalid, it will return a 400 error 
	// to the client and tell them which fields are invalid.
	if err:= c.ShouldBindJSON(&input); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	book := models.Book{Title: input.Title, Author: input.Author}
	models.DB.Create(&book)

	c.JSON(http.StatusOK, gin.H{
		"data": book,
	})
}

func FindBook(c *gin.Context){
	var book models.Book

	if err := models.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error":"not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data":book})
}

func UpdateBook(c *gin.Context) {
	// Get model if exist
	var book models.Book
	if err := models.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	fmt.Println("select book:", book)
	// Validate input
	var input models.UpdateBookInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println("user input:", input)
	uperr:=models.DB.Model(&book).Updates(models.Book{Title:input.Title, Author:input.Author}).Error
	if uperr != nil{
		panic(uperr)
	}
	c.JSON(http.StatusOK, gin.H{"data": book})
}

func DeleteBook(c *gin.Context){
	var book models.Book
	fmt.Println(book)
	if err := models.DB.Where("id = ? ", c.Param("id")).First(&book).Error; err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": "not found"})
	}

	models.DB.Delete(&book)
	c.JSON(http.StatusOK, gin.H{"result": "delete complete"})
}