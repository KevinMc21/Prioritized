package tasks

type TaskCategory struct{
	Name		string		`json:"name"`
	Difficulty	float32		`json:"difficulty"`
	TimeRanges	[]TimePeriod	`json:"time_ranges"`
}