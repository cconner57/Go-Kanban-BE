package routes

import (
	"kanban/api/types"
	"net/http"

	"github.com/gin-gonic/gin"
)

var tasks = []types.Task{
	{ID: "1", Title: "Test Title 1", Details: "Test Details 1"},
	{ID: "2", Title: "Test Title 2", Details: "Test Details 2"},
	{ID: "3", Title: "Test Title 3", Details: "Test Details 3"},
}

func GetTasks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, tasks)
}