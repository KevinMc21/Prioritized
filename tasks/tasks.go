package tasks

import (
	"encoding/json"
	"time"
)

// Structure of task data
type Task struct{
	Name		string		`json:"name,omitempty" validate:"required"`
	Timeline	Period		`json:"timeline,omitempty"` // Provide a start time and deadline for a task. Optional
	AssignedTime	Period		`json:"assigned_time,omitempty"` // time assigned to task
	Fixed		bool		`json:"fixed,omitempty"`
	EstimatedTime	time.Duration	`json:"estimated_time,omitempty" validate:"required,min=1m"` // How much time to spend on the task
	WeightCoef	float64		`json:"weight_coef,omitempty" validate:"required,min=1,max=2"` // The weight coefficient for the specific task. Grouping coefficient will also be considered 
	CurrentScore	float64		`json:"current_score"`
}

func (t *Task) UnmarshalJSON(b []byte) error {
	var temp struct{
		Name		string		`json:"name,omitempty" validate:"required"`
		Timeline	Period	`json:"timeline,omitempty"` // Provide a start time and deadline for a task. Optional
		AssignedTime	Period	`json:"assigned_time,omitempty"` // time assigned to task
		Fixed		bool		`json:"fixed,omitempty"`
		EstimatedTime	string		`json:"estimated_time,omitempty" validate:"required,min=1m"` // How much time to spend on the task
		WeightCoef	float64		`json:"weight_coef,omitempty" validate:"required,min=1,max=2"` // The weight coefficient for the specific task. Grouping coefficient will also be considered 
		CurrentScore	float64		`json:"current_score"`
	}

	var err error
	if err = json.Unmarshal(b, &temp); err != nil {
		return err
	}

	t.Name = temp.Name
	t.Timeline = temp.Timeline
	t.AssignedTime = temp.AssignedTime
	t.Fixed = temp.Fixed
	t.EstimatedTime, err = time.ParseDuration(temp.EstimatedTime)
	t.WeightCoef = temp.WeightCoef
	t.CurrentScore = temp.CurrentScore

	return err
}

func (t *Task) MarshalJSON() ([]byte, error) {
	var temp struct{
		Name		string		`json:"name,omitempty" validate:"required"`
		Timeline	Period		`json:"timeline,omitempty"` // Provide a start time and deadline for a task. Optional
		AssignedTime	Period		`json:"assigned_time,omitempty"` // time assigned to task
		Fixed		bool		`json:"fixed,omitempty"`
		EstimatedTime	string		`json:"estimated_time,omitempty" validate:"required,min=1m"` // How much time to spend on the task
		WeightCoef	float64		`json:"weight_coef,omitempty"` // The weight coefficient for the specific task. Grouping coefficient will also be considered 
		CurrentScore	float64		`json:"current_score"`
	}

	temp.Name = t.Name
	temp.Timeline = t.Timeline
	temp.AssignedTime = t.AssignedTime
	temp.Fixed = t.Fixed
	temp.EstimatedTime = t.EstimatedTime.String()
	temp.WeightCoef = t.WeightCoef
	temp.CurrentScore = t.CurrentScore

	json, err := json.Marshal(temp)
	if err != nil { return nil, err }

	return json, nil
}


