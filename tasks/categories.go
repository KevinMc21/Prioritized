package tasks

// Categories to seperate different tasks into, ParentCategory field specifies what category
// it is a sub-category of (if any)
type TaskCategory struct{
	ID		int		`json:"id"`
	Name		string		`json:"name"`
	ParentCategory	int		`json:"sub_category"`
	CostRating	float64		`json:"cost_rating"`
	TimeRanges	[]TimePeriod	`json:"time_ranges"`
}


