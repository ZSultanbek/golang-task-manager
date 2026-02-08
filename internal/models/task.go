package models
/*
Task
- id (int)
- Title (string)
- done (bool)
*/

type Task struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}


