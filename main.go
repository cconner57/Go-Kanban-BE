package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"

	R "kanban/api/handler"
)

func main() {
	r := chi.NewRouter()
	
	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"access-control-allow-headers", "Content-Type", "access-control-allow-origin"},
	})

	r.Use(cors.Handler)
	r.Get("/", R.GetTasks)


	http.ListenAndServe(":3000", r)
}
