package routes

import (
	"kanban/api/types"
	"net/http"
	"github.com/gin-gonic/gin"
)

var tasks = []types.Task{
	{ID: "1", Title: "Build UI for onboarding flow", Column: "todo", SubTask: []types.SubTaskItem{{Finished: true, Description: ""},{Finished: true, Description: ""}}},
	{ID: "2", Title: "Build UI for search", Column: "todo", SubTask: []types.SubTaskItem{{Finished: true, Description: ""},{Finished: true, Description: ""}}},
	{ID: "3", Title: "Build settings UI", Column: "todo", SubTask: []types.SubTaskItem{{Finished: true, Description: ""},{Finished: true, Description: ""}}},
	{ID: "3", Title: "QA and test all major user journeys", Column: "todo", SubTask: []types.SubTaskItem{{Finished: true, Description: ""},{Finished: true, Description: ""}}},
}

func GetTasks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, tasks)
}