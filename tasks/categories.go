package tasks

// Categories to seperate different tasks into, ParentCategory field specifies what category
// it is a sub-category of (if any)
type TaskGrouping struct{
	Name		string		`json:"name"`
	WeightCoef	float64		`json:"weight_coef" validate:"required,min=1,max=2"` // The weigh coefficient that will be used to scale the task score
	TimeRanges	[]timePeriod	`json:"time_ranges" validate:"required,min=1"` // Specifies the time range that the tasks in the group can be done. I.e. Work can be done from 9a.m - 5p.m
	Tasks		[]Task		`json:"tasks"` // The list of the users tasks
}

type currentTasks struct{
	AssignedTime	timePeriod	`json:"assigned_time"`
	Timeline	timePeriod	`json:"timeline"`
	Fixed		bool		`json:"fixed"`
}

