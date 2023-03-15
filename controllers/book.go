package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tabeyla/rest-grom-postgres/models"
)

func FindBooks(c *gin.Context) {

	var books []models.Book
	models.DB.Find(&books)
	c.JSON(http.StatusOK, gin.H{"data": books})
}

func FindBook(c *gin.Context) {

	var book models.Book
	if err := models.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": book})
}

func CreateBook(c *gin.Context) {

	var input models.CreateBookInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return

	}

	book := models.Book{Title: input.Title, Author: input.Author}

	models.DB.Create(&book)
	c.JSON(http.StatusOK, gin.H{"data": book})

}

func UpdateBook(c *gin.Context) {

	var input models.UpdateBookInput
	//check input
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//check if record exists
	var book models.Book

	if err := models.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}

	//update record if exists
	models.DB.Model(&book).Update(&input)

	c.JSON(http.StatusOK, gin.H{"data": book})

}

func DeleteBook(c *gin.Context) {

	var book models.Book
	if err := models.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	models.DB.Delete(&book)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
