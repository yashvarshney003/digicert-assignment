package controllers

import (
	"fmt"
	"public_library/database"
	"public_library/models"
	"strconv"
)

var GetAllBooks = func() ([]models.Book, error) {
	var books []models.Book
	result := database.DB.Find(&books)
	return books, result.Error
}

var GetBookByID = func(id string) (*models.Book, error) {
	var book models.Book
	result := database.DB.First(&book, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &book, nil
}

var CreateBookDal = func(book *models.Book) (*models.Book, error) {
	result := database.DB.Create(book)
	if result.Error != nil {
		return nil, result.Error
	}
	return book, nil
}

var UpdateBookDal = func(id uint, input *BookUpdateInput) (*models.Book, error) {
	book, err := GetBookByID(strconv.FormatUint(uint64(id), 10))
	if err != nil {
		return nil, err
	}

	book.Title = input.Title
	book.Description = input.Description
	book.Author = input.Author

	result := database.DB.Save(book)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, fmt.Errorf("no rows updated")
	}

	return book, nil
}

var DeleteBookByID = func(id string) error {
	book, err := GetBookByID(id)
	if err != nil {
		return err
	}

	result := database.DB.Delete(book)
	return result.Error
}
