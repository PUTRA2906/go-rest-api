package router

import (
	"net/http"
	"restapi/controller"

	"github.com/gorilla/mux"
)

func NewRouter(c *controller.BookController) http.Handler {
    r := mux.NewRouter()

    r.HandleFunc("/books", c.GetBooks).Methods("GET")
    r.HandleFunc("/books/{id}", c.GetBook).Methods("GET")
    r.HandleFunc("/books", c.CreateBook).Methods("POST")
    r.HandleFunc("/books/{id}", c.UpdateBook).Methods("PUT")
    r.HandleFunc("/books/{id}", c.DeleteBook).Methods("DELETE")

    return r
}
