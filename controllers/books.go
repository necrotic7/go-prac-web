package controllers

import (
	"gin-bookstore/m/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func FindBooks(c *gin.Context){
	var books []models.Book
	models.DB.Find(&books)

	c.JSON(http.StatusOK, gin.H{"data": books})
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