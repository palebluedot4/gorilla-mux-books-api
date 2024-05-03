package controller

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"

	"gorilla-mux-books-api/cmd/controller/model"
)

var log = logrus.New()

func init() {
	log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
	log.SetLevel(logrus.InfoLevel)
}

func GetBooksHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(model.Books); err != nil {
		log.WithError(err).Error("Error encoding books")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func GetBookByIDHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)

	books := model.Books
	for _, v := range books {
		if v.ID == vars["id"] {
			if err := json.NewEncoder(w).Encode(v); err != nil {
				log.WithError(err).Error("Error encoding book")
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			return
		}
	}

	http.Error(w, "Book not found", http.StatusNotFound)
}

func CreateBookHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var book model.Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		log.WithError(err).Error("Error decoding book")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	model.Books = append(model.Books, &book)

	if err := json.NewEncoder(w).Encode(book); err != nil {
		log.WithError(err).Error("Error encoding book")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func DeleteBookHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)

	var found bool
	for i, v := range model.Books {
		if v.ID == vars["id"] {
			model.Books = append(model.Books[:i], model.Books[i+1:]...)
			found = true
			break
		}
	}

	if !found {
		http.Error(w, "Book not found: "+vars["id"], http.StatusNotFound)
		return
	}

	if err := json.NewEncoder(w).Encode(model.Books); err != nil {
		log.WithError(err).Error("Error encoding books")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func UpdateBookHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)

	updatedBook := model.Book{}
	if err := json.NewDecoder(r.Body).Decode(&updatedBook); err != nil {
		log.WithError(err).Error("Error decoding updated book")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	for i, v := range model.Books {
		if v.ID == vars["id"] {
			if updatedBook.Title != "" {
				model.Books[i].Title = updatedBook.Title
			}
			if updatedBook.ISBN != "" {
				model.Books[i].ISBN = updatedBook.ISBN
			}
			if updatedBook.Author != nil {
				model.Books[i].Author = updatedBook.Author
			}

			if err := json.NewEncoder(w).Encode(model.Books[i]); err != nil {
				log.WithError(err).Error("Error encoding updated book")
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			return
		}
	}

	http.Error(w, "Book not found: "+vars["id"], http.StatusNotFound)
}
