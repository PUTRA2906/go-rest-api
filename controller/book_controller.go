package controller

import (
	"encoding/json"
	"net/http"
	"restapi/model"
	"restapi/service"
	"strconv"

	"github.com/gorilla/mux"
)

type BookController struct {
    Service service.BookService
}

func NewBookController(service service.BookService) *BookController {
    return &BookController{Service: service}
}

func (c *BookController) GetBooks(w http.ResponseWriter, r *http.Request) {
    json.NewEncoder(w).Encode(c.Service.GetBooks())
}

func (c *BookController) GetBook(w http.ResponseWriter, r *http.Request) {
    id, _ := strconv.Atoi(mux.Vars(r)["id"])
    book, found := c.Service.GetBook(id)
    if !found {
        http.NotFound(w, r)
        return
    }
    json.NewEncoder(w).Encode(book)
}

func (c *BookController) CreateBook(w http.ResponseWriter, r *http.Request) {
    var book model.Book
    json.NewDecoder(r.Body).Decode(&book)
    json.NewEncoder(w).Encode(c.Service.CreateBook(book))
}

func (c *BookController) UpdateBook(w http.ResponseWriter, r *http.Request) {
    id, _ := strconv.Atoi(mux.Vars(r)["id"])
    var book model.Book
    json.NewDecoder(r.Body).Decode(&book)
    updated, ok := c.Service.UpdateBook(id, book)
    if !ok {
        http.NotFound(w, r)
        return
    }
    json.NewEncoder(w).Encode(updated)
}

func (c *BookController) DeleteBook(w http.ResponseWriter, r *http.Request) {
    id, _ := strconv.Atoi(mux.Vars(r)["id"])
    ok := c.Service.DeleteBook(id)
    if !ok {
        http.NotFound(w, r)
        return
    }
    w.Write([]byte("Buku berhasil dihapus"))
}
