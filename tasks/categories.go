package tasks

// Categories to seperate different tasks into, ParentCategory field specifies what category
// it is a sub-category of (if any)
type TaskGrouping struct{
	ID		int		`json:"id"` // User specific ID for task grouping
	Name		string		`json:"name"` // Grouping name. I.e. Work, personal, school, etc
	ChildOf		int		`json:"child_of"` // Specifies the parent grouping of the task based on ID
	WeightCoef	float64		`json:"weight_coef"` // The weigh coefficient that will be used to scale the task score
	TimeRanges	[]TimePeriod	`json:"time_ranges"` // Specifies the time range that the tasks in the group can be done. I.e. Work can be done from 9a.m - 5p.m
	UserTasks	[]UserTask	`json:"user_tasks"` // The list of the users tasks
}


