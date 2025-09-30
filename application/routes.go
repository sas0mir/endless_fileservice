package application

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/sas0mir/endless_fileservice/handler"
)

func loadRoutes() *chi.Mux {
	router := chi.NewRouter()

	router.Use(middleware.Logger)

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	router.Route("/file", loadFileRoutes)

	return router
}

func loadFileRoutes(router chi.Router) {
	fileHandler := &handler.File{}

	router.Post("/upload", fileHandler.Upload)
	router.Get("/list", fileHandler.List)
	router.Get("/download/{fileID}", fileHandler.Download)
	router.Delete("/delete/{fileID}", fileHandler.Delete)
}
