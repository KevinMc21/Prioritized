package server

type Operation int

const (
	Insert Operation = iota // Inserts a task
	Delete			// Deletes a task
	ReSort			// Provided groupings' tasks gets resorted
	Update			// Edits a task or task grouping
	Anchor			// Sets the fixed attribute for task or task grouping to be true
	Move			// Move a task to a specific time
)