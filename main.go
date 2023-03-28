package main

import (
	"net/http"

	R "kanban/api/handler"
	C "kanban/api/logic"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"

	_ "github.com/lib/pq"
)



func main() {
	C.ConnectDB()

	r := chi.NewRouter()
	
	cors := cors.New(cors.Options{
		AllowedOrigins:	[]string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept, Content-Type, Content-Length, Authorization"},
		AllowCredentials: true,
	})

	r.Use(cors.Handler)
	r.Get("/", R.GetTasks)

	http.ListenAndServe(":3000", r)
}