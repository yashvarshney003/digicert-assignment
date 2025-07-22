package controllers

import (
	"errors"
	"net/http"
	"public_library/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetBooks(c *gin.Context) {
	result, err := GetAllBooks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	books_output := make([]BookOutput, 0)
	for _, book := range result {
		book_out := BookOutput{
			ID:          book.ID,
			Title:       book.Title,
			Description: book.Description,
			Author:      book.Author,
		}
		books_output = append(books_output, book_out)
	}
	c.JSON(http.StatusOK, books_output)
}

func GetBook(c *gin.Context) {
	id := c.Param("id")
	book, err := GetBookByID(id)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}
	book_output := BookOutput{
		ID:          book.ID,
		Title:       book.Title,
		Description: book.Description,
		Author:      book.Author,
	}
	c.JSON(http.StatusOK, book_output)
}

func CreateBook(c *gin.Context) {
	var input BookInput
	var output BookOutput

	if err1 := c.ShouldBindJSON(&input); err1 != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err1.Error()})
		return
	}
	book := models.Book{
		Title:       input.Title,
		Description: input.Description,
		Author:      input.Author,
	}
	result, err := CreateBookDal(&book)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	output = BookOutput{
		ID:          result.ID,
		Title:       result.Title,
		Description: result.Description,
		Author:      result.Author,
	}

	c.JSON(http.StatusCreated, output)
}

func UpdateBook(c *gin.Context) {

	var input BookUpdateInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	book, err := UpdateBookDal(input.ID, &input)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	book_output := BookOutput{
		ID:          book.ID,
		Title:       book.Title,
		Description: book.Description,
		Author:      book.Author,
	}

	c.JSON(http.StatusOK, book_output)
}

func DeleteBook(c *gin.Context) {
	id := c.Param("id")
	err := DeleteBookByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Book deleted successfully"})
}
