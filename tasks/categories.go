package tasks

// Categories to seperate different tasks into, SubCategory field specifies what category
// it is a sub-category of (if any)
type TaskCategory struct{
	Name		string		`json:"name"`
	Difficulty	float32		`json:"difficulty"`
	SubCategory	*TaskCategory	`json:"sub_category"`
	TimeRanges	[]TimePeriod	`json:"time_ranges"`
}