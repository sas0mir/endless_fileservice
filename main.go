package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {

	router := chi.NewRouter()

	router.Get("/files", basicHandler)

	server := &http.Server{
		Addr:    ":3002",
		Handler: router,
	}

	err := server.ListenAndServe()
	if err != nil {
		fmt.Println("Failed to listen to server ", err)
	}
}

func basicHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		if r.URL.Path == "/foo" {
		}
	}

	if r.Method == http.MethodPost {
	}

	w.Write([]byte("Hello, World!"))
}
