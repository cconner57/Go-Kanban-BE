package main

import (
	"net/http"

	R "kanban/api/handler"
	C "kanban/api/logic"

	"github.com/go-chi/chi/v5"
	// "github.com/go-chi/cors"

	_ "github.com/lib/pq"
)



func main() {
	C.ConnectDB()

	r := chi.NewRouter()
	
	// cors := cors.New(cors.Options{
	// 	AllowedOrigins:   []string{"*"},
	// 	// AllowCredentials: true,
	// 	// AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
	// 	AllowedHeaders:   []string{"access-control-allow-headers", "Content-Type", "access-control-allow-origin", "x-requested-with", "Origin", "Accept"},
	// })

	// r.Use(cors.Handler)
	r.Get("/", R.GetTasks)

	http.ListenAndServe(":3000", r)
}