package tasks

import "time"

// Structure of task data
type Task struct{
	Name		string		`json:"name"`
	Description 	string		`json:"description"`
	Category	string		`json:"category"`
	Timeline	DatetimePeriod	`json:"timeline"`
	EstimatedTime	time.Duration	`json:"estimated_time"`
	TaskDifficulty	float32		`json:"task_difficulty"`
	Score		int32		`json:"score"`
}
