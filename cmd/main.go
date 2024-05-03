package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"

	"gorilla-mux-books-api/cmd/controller"
	"gorilla-mux-books-api/cmd/controller/model"
)

var log = logrus.New()

func init() {
	log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
	log.SetLevel(logrus.InfoLevel)
}

func main() {
	model.Books = append(model.Books, &model.Book{ID: "1", ISBN: "9789574157327", Title: "腹語術", Author: &model.Author{FirstName: "宇", LastName: "夏"}})
	model.Books = append(model.Books, &model.Book{ID: "2", ISBN: "9789571375922", Title: "大裂", Author: &model.Author{FirstName: "遷", LastName: "胡"}})

	router := mux.NewRouter()
	router.HandleFunc("/books", controller.GetBooksHandler).Methods("GET")
	router.HandleFunc("/books/{id}", controller.GetBookByIDHandler).Methods("GET")
	router.HandleFunc("/books", controller.CreateBookHandler).Methods("POST")
	router.HandleFunc("/books/{id}", controller.DeleteBookHandler).Methods("DELETE")
	router.HandleFunc("/books/{id}", controller.UpdateBookHandler).Methods("PATCH")

	log.Info("Listening and serving HTTP on localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
