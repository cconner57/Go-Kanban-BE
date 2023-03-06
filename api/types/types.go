package types

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