package models

import "time"

// TODO holds the data of a todo
type TODO struct {
	Title       string
	Content     string
	IsImportant bool
	IsDone      bool
	Time        time.Time
}
