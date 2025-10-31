package handlers

import (
	"encoding/json"
	"net/http"
	"server/cmd/models"
	"server/cmd/repositories"
	"strconv"
	"strings"
)

type BookHandler struct {
	repo repositories.BookRepository
}

func NewBookHandler(repo repositories.BookRepository) *BookHandler {
	return &BookHandler{repo: repo}
}

func (h *BookHandler) HandleBooks(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	if path == "/books/" {
		switch r.Method {
		case http.MethodGet:
			h.getAllBooks(w, r)
		case http.MethodPost:
			h.createBook(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	} else {
		parts := strings.Split(path, "/")
		if len(parts) != 3 || parts[1] != "books" {
			http.Error(w, "Not found", http.StatusNotFound)
			return
		}
		id, err := strconv.Atoi(parts[2])
		if err != nil {
			http.Error(w, "Invalid book ID", http.StatusBadRequest)
			return
		}
		switch r.Method {
		case http.MethodGet:
			h.getBook(w, r, id)
		case http.MethodPut:
			h.updateBook(w, r, id)
		case http.MethodDelete:
			h.deleteBook(w, id)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}

}

func (h *BookHandler) getAllBooks(w http.ResponseWriter, r *http.Request) {
	books, err := h.repo.GetAll()
	if err != nil {
		http.Error(w, "Failed to retrieve books", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}
func (h *BookHandler) getBook(w http.ResponseWriter, r *http.Request, id int) {
	book, err := h.repo.GetByID(id)
	if err != nil {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(book)
}

func (h *BookHandler) createBook(w http.ResponseWriter, r *http.Request) {
	var book models.Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	createdBook, err := h.repo.Create(&book)
	if err != nil {
		http.Error(w, "Failed to create book", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdBook)
}

func (h *BookHandler) updateBook(w http.ResponseWriter, r *http.Request, id int) {
	var book models.Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	updatedBook, err := h.repo.Update(id, book)
	if err != nil {
		http.Error(w, "Failed to update book", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedBook)
}

func (h *BookHandler) deleteBook(w http.ResponseWriter, id int) {
	if err := h.repo.Delete(id); err != nil {
		http.Error(w, "Failed to delete book", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
