package routes

import (
	"net/http"
	"server/cmd/handlers"
)

func SetupBookRoutes(bookHandler *handlers.BookHandler) {
	http.HandleFunc("/books/", bookHandler.HandleBooks)
}
