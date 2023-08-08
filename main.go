package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello World!"))
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/hello", helloHandler)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}
	log.Println("Server is listening on port 8080")
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
