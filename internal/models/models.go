package models

import "time"

// TODO holds the data of a todo
type TODO struct {
	Title  string
	Time   time.Time
	IsDone bool
}
