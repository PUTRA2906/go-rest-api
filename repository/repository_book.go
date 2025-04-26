package repository

import (
	"restapi/config"
	"restapi/model"
)

type BookRepository interface {
    FindAll() []model.Book
    FindByID(id int) (model.Book, bool)
    Save(book model.Book) model.Book
    Update(id int, book model.Book) (model.Book, bool)
    Delete(id int) bool
}

type bookRepo struct{}

func NewBookRepository() BookRepository {
    return &bookRepo{}
}

func (r *bookRepo) FindAll() []model.Book {
    var books []model.Book
    config.DB.Find(&books)
    return books
}

func (r *bookRepo) FindByID(id int) (model.Book, bool) {
    var book model.Book
    result := config.DB.First(&book, id)
    return book, result.Error == nil
}

func (r *bookRepo) Save(book model.Book) model.Book {
    config.DB.Create(&book)
    return book
}

func (r *bookRepo) Update(id int, updated model.Book) (model.Book, bool) {
    var book model.Book
    result := config.DB.First(&book, id)
    if result.Error != nil {
        return model.Book{}, false
    }
    book.Title = updated.Title
    book.Author = updated.Author
    config.DB.Save(&book)
    return book, true
}

func (r *bookRepo) Delete(id int) bool {
    result := config.DB.Delete(&model.Book{}, id)
    return result.RowsAffected > 0
}
