package types

type Board struct {
	ID string `json:"id"`
	Title string `json:"title"`
	Column []Column `json:"column"`
}

type Column struct {
	ID string `json:"id"`
	Title string `json:"title"`
	Task []Task `json:"task"`
}

type Task struct {
	ID string `json:"id"`
	Title string `json:"title"`
	Column string `json:"column"`
	SubTask []SubTaskItem `json:"subTask"`
}

type SubTaskItem struct {
	Finished bool `json:"finished"`
	Description string `json:"description"`
}