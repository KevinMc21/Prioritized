package tasks

import "time"

// Structure of task data
type Task struct{
	Name		string		`json:"name"`
	Description 	string		`json:"description"`
	Category	string		`json:"category"`
	Timeline	timePeriod	`json:"timeline"`
	EstimatedTime	time.Duration	`json:"estimated_time"`
	TaskDifficulty	float32		`json:"task_difficulty"`
}

type timePeriod struct{
	TimeStart 	time.Time	`json:"time_start"`
	TimeEnd		time.Time	`json:"time_end"`
}