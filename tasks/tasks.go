package tasks

import "time"

// Structure of task data
type Task struct{
	ID		int		`json:"id"`
	Name		string		`json:"name"`
	Description 	string		`json:"description"`
	Category	string		`json:"category"`
	Timeline	DatetimePeriod	`json:"timeline"`
	EstimatedTime	time.Duration	`json:"estimated_time"`
	TaskRating	float64		`json:"task_rating"`
	Score		int		`json:"score"`
}

