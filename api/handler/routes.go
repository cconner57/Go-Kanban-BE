package routes

import (
	"encoding/json"
	"kanban/api/types"
	"net/http"
)

var board = []types.Board{
	{ID: "1", Title: "Platform Launch", Column: []types.Column{
		{ID: "1", Title: "To Do", Color: "#49c4e5", Task: []types.Task{
			{ID: "1", Title: "Build UI for onboarding flow", Column: "todo", SubTask: []types.SubTaskItem{{Finished: true, Description: ""},{Finished: true, Description: ""}}},
			{ID: "2", Title: "Build UI for search", Column: "todo", SubTask: []types.SubTaskItem{{Finished: true, Description: ""},{Finished: true, Description: ""}}},
			{ID: "3", Title: "Build settings UI", Column: "todo", SubTask: []types.SubTaskItem{{Finished: true, Description: ""},{Finished: true, Description: ""}}},
		}},
		{ID: "2", Title: "In Progress", Color: "#8471F2", Task: []types.Task{
			{ID: "4", Title: "Build settings UI", Column: "todo", SubTask: []types.SubTaskItem{{Finished: true, Description: ""},{Finished: true, Description: ""}}},
		}},
		{ID: "3", Title: "Done", Color: "#67E2AE", Task: []types.Task{
			{ID: "5", Title: "QA and test all major user journeys", Column: "todo", SubTask: []types.SubTaskItem{{Finished: true, Description: ""},{Finished: true, Description: ""}}},
		}},
	}},
	{ID: "2", Title: "Marketing Plan", Column: []types.Column{
		{ID: "1", Title: "To Do", Color: "#49c4e5", Task: []types.Task{
			{ID: "1", Title: "Build UI for onboarding flow", Column: "todo", SubTask: []types.SubTaskItem{{Finished: true, Description: ""},{Finished: true, Description: ""}}},
			{ID: "2", Title: "Build UI for search", Column: "todo", SubTask: []types.SubTaskItem{{Finished: true, Description: ""},{Finished: true, Description: ""}}},
		}},
		{ID: "2", Title: "In Progress", Color: "#8471F2", Task: []types.Task{
			{ID: "4", Title: "Build settings UI", Column: "todo", SubTask: []types.SubTaskItem{{Finished: true, Description: ""},{Finished: true, Description: ""}}},
		}},
		{ID: "3", Title: "Done", Color: "#67E2AE", Task: []types.Task{
			{ID: "5", Title: "QA and test all major user journeys", Column: "todo", SubTask: []types.SubTaskItem{{Finished: true, Description: ""},{Finished: true, Description: ""}}},
		}},
	}},
	{ID: "3", Title: "Roadmap", Column: []types.Column{
		
		{ID: "2", Title: "To Do", Color: "#8471F2", Task: []types.Task{
			{ID: "4", Title: "Build settings UI", Column: "todo", SubTask: []types.SubTaskItem{{Finished: true, Description: ""},{Finished: true, Description: ""}}},
		}},
		{ID: "3", Title: "In Progress", Color: "#67E2AE", Task: []types.Task{
			{ID: "5", Title: "QA and test all major user journeys", Column: "todo", SubTask: []types.SubTaskItem{{Finished: true, Description: ""},{Finished: true, Description: ""}}},
		}},
		{ID: "1", Title: "Done", Color: "#49c4e5", Task: []types.Task{
			{ID: "1", Title: "Build UI for onboarding flow", Column: "todo", SubTask: []types.SubTaskItem{{Finished: true, Description: ""},{Finished: true, Description: ""}}},
			{ID: "2", Title: "Build UI for search", Column: "todo", SubTask: []types.SubTaskItem{{Finished: true, Description: ""},{Finished: true, Description: ""}}},
			{ID: "3", Title: "Build settings UI", Column: "todo", SubTask: []types.SubTaskItem{{Finished: true, Description: ""},{Finished: true, Description: ""}}},
		}},
	}},
}

func GetTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(board)	
}