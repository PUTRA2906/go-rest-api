package service

import (
	"restapi/model"
	"restapi/repository"
)

type BookService interface {
    GetBooks() []model.Book
    GetBook(id int) (model.Book, bool)
    CreateBook(book model.Book) model.Book
    UpdateBook(id int, book model.Book) (model.Book, bool)
    DeleteBook(id int) bool
}

type bookService struct {
    repo repository.BookRepository
}

func NewBookService(repo repository.BookRepository) BookService {
    return &bookService{repo: repo}
}

func (s *bookService) GetBooks() []model.Book {
    return s.repo.FindAll()
}

func (s *bookService) GetBook(id int) (model.Book, bool) {
    return s.repo.FindByID(id)
}

func (s *bookService) CreateBook(book model.Book) model.Book {
    return s.repo.Save(book)
}

func (s *bookService) UpdateBook(id int, book model.Book) (model.Book, bool) {
    return s.repo.Update(id, book)
}

func (s *bookService) DeleteBook(id int) bool {
    return s.repo.Delete(id)
}
