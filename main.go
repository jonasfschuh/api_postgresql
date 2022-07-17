package main

import (
	"fmt"
	"net/http"

	"github.com/WilkerAlves/api-postgresql/configs"
	"github.com/WilkerAlves/api-postgresql/handlers"
	"github.com/go-chi/chi/v5"
)

func main() {

	fmt.Println("teste")
	err := configs.Load()
	if err != nil {
		panic(err)
	}

	r := chi.NewRouter()
	r.Post("/", handlers.Create)
	r.Put("/{id}", handlers.Update)
	r.Delete("/{id}", handlers.Delete)
	r.Get("/{id}", handlers.Get)
	r.Get("/", handlers.List)

	http.ListenAndServe(fmt.Sprintf(":%s", configs.GetServerPort()), r)
}
