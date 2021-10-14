package tasks

import (
	"encoding/json"
	"time"
)

// Structure of task data
type Task struct{
	Timeline	timePeriod	`json:"timeline"` // Provide a start time and deadline for a task. Optional
	AssignedTime	timePeriod	`json:"assigned_time"` // time assigned to task
	Fixed		bool		`json:"fixed"`
	EstimatedTime	time.Duration	`json:"estimated_time" validate:"required,min=1m"` // How much time to spend on the task
	WeightCoef	float64		`json:"weight_coef" validate:"required,min=1,max=2"` // The weight coefficient for the specific task. Grouping coefficient will also be considered 
	CurrentScore	float64		`json:"current_score"`
}

func (t *Task) UnmarshalJSON(b []byte) error {
	var temp struct{
		Timeline	timePeriod	`json:"timeline"` // Provide a start time and deadline for a task. Optional
		AssignedTime	timePeriod	`json:"assigned_time"` // time assigned to task
		Fixed		bool		`json:"fixed"`
		EstimatedTime	string		`json:"estimated_time" validate:"required,min=1m"` // How much time to spend on the task
		WeightCoef	float64		`json:"weight_coef" validate:"required,min=1,max=2"` // The weight coefficient for the specific task. Grouping coefficient will also be considered 
		CurrentScore	float64		`json:"current_score"`
	}

	var err error
	if err = json.Unmarshal(b, &temp); err != nil {
		return err
	}

	t.Timeline = temp.Timeline
	t.AssignedTime = temp.AssignedTime
	t.Fixed = temp.Fixed
	t.EstimatedTime, err = time.ParseDuration(temp.EstimatedTime)
	t.WeightCoef = temp.WeightCoef
	t.CurrentScore = temp.CurrentScore

	return err
}


