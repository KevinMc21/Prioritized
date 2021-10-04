package tasks

import "time"

// Structure of task data
type Task struct{
	Name		string		`json:"name"` // Task Name
	Description 	string		`json:"description"` // Description for the task
	Timeline	DatetimePeriod	`json:"timeline"` // Provide a start time and deadline for a task. Optional
	EstimatedTime	time.Duration	`json:"estimated_time"` // How much time to spend on the task
	WeightCoef	float64		`json:"task_rating"` // The weight coefficient for the specific task. Grouping coefficient will also be considered 
}

type UserTask struct {
	ID		int		`json:"id"`
	Task		Task
	AssignedTime	DatetimePeriod
	Fixed		bool
	Score		float64
}

