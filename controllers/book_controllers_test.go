package controllers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"public_library/models"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func NewMockBook(id uint, title, desc, author string) models.Book {
	return models.Book{
		Model:       gorm.Model{ID: id},
		Title:       title,
		Description: desc,
		Author:      author,
	}
}

var (
	mockBooks = []models.Book{
		NewMockBook(1, "Book 1", "Desc 1", "Author 1"),
	}

	mockBook = NewMockBook(1, "Book 1", "Desc 1", "Author 1")
)

func mockGetAllBooks() ([]models.Book, error) {
	return mockBooks, nil
}

func mockGetBookByID(id string) (*models.Book, error) {
	if id == "1" {
		return &mockBook, nil
	}
	return nil, gorm.ErrRecordNotFound
}

func mockCreateBookDal(book *models.Book) (*models.Book, error) {
	book.ID = 2
	return book, nil
}

func mockUpdateBookDal(id uint, input *BookUpdateInput) (*models.Book, error) {
	if id != 1 {
		return nil, gorm.ErrRecordNotFound
	}
	return &mockBook, nil
}

func mockDeleteBookByID(id string) error {
	if id == "1" {
		return nil
	}
	return gorm.ErrRecordNotFound
}

func mockGetAllBooksEmpty() ([]models.Book, error) {
	return []models.Book{}, nil
}

func init() {
	GetAllBooks = mockGetAllBooks
	GetBookByID = mockGetBookByID
	CreateBookDal = mockCreateBookDal
	UpdateBookDal = mockUpdateBookDal
	DeleteBookByID = mockDeleteBookByID
}

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/books", GetBooks)
	r.GET("/books/:id", GetBook)
	r.POST("/books", CreateBook)
	r.PUT("/books", UpdateBook)
	r.DELETE("/books/:id", DeleteBook)
	return r
}

func TestGetBooks(t *testing.T) {
	router := setupRouter()
	req, _ := http.NewRequest("GET", "/books", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	expected_body := `[{"id":1,"title": "Book 1","description":"Desc 1","author":"Author 1"}]`
	assert.Equal(t, http.StatusOK, w.Code)
	assert.JSONEq(t, expected_body, w.Body.String())
}

func TestGetBooks_NoRecords(t *testing.T) {

	originalGetAllBooks := GetAllBooks
	GetAllBooks = mockGetAllBooksEmpty
	defer func() {
		GetAllBooks = originalGetAllBooks
	}()

	router := setupRouter()
	req, _ := http.NewRequest("GET", "/books", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)

}
func TestGetBook_Success(t *testing.T) {
	router := setupRouter()
	req, _ := http.NewRequest("GET", "/books/1", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	expected_body := `{"id":1,"title": "Book 1","description":"Desc 1","author":"Author 1"}`
	assert.Equal(t, http.StatusOK, w.Code)
	assert.JSONEq(t, expected_body, w.Body.String())
}

func TestGetBook_NotFound(t *testing.T) {
	router := setupRouter()
	req, _ := http.NewRequest("GET", "/books/999", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
	assert.Contains(t, w.Body.String(), "Book not found")
}

func TestCreateBook(t *testing.T) {
	router := setupRouter()
	body := `{"title": "New Book", "description": "Desc", "author": "Author"}`
	req, _ := http.NewRequest("POST", "/books", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Contains(t, w.Body.String(), "New Book")
}

func TestUpdateBook_BadRequestMissingRequiredFields(t *testing.T) {
	router := setupRouter()
	body := `{"id": 1, "title": "Updated Title"}`
	req, _ := http.NewRequest("PUT", "/books", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "Error:Field validation for 'Author'")
}

func TestUpdateBook_NotFound(t *testing.T) {
	router := setupRouter()
	body := `{"id": 999, "title": "Unknown"}`
	req, _ := http.NewRequest("PUT", "/books", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestDeleteBook_Success(t *testing.T) {
	router := setupRouter()
	req, _ := http.NewRequest("DELETE", "/books/1", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Book deleted successfully")
}

func TestDeleteBook_NotFound(t *testing.T) {
	router := setupRouter()
	req, _ := http.NewRequest("DELETE", "/books/999", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
}
