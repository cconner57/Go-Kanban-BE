package main

import (
	"github.com/gin-gonic/gin"
	"kanban/api/handler"
)

func main() {
	router := gin.Default()

	router.GET("/", routes.GetTasks)

	router.Run()
}
