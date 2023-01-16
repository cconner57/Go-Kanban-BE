package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	"kanban/api/handler"
)

func main() {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
        AllowOrigins: []string{"*"},
        AllowMethods: []string{"POST", "PUT", "PATCH", "DELETE"},
        AllowHeaders: []string{"Content-Type,access-control-allow-origin, access-control-allow-headers"},
    }))

	router.GET("/", routes.GetTasks)

	router.Run()
}
