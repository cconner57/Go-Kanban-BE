package types

type Board struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Column []Column `json:"column"`
}

type Column struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Color string `json:"color"`
	Task []Task `json:"task"`
}

type Task struct {
	Id string `json:"id"`
	Name string `json:"name"`
	SubTask []SubTask `json:"subTask"`
}

type SubTask struct {
	Completed bool `json:"completed"`
	Description string `json:"description"`
}