package dbModels

import "time"

type User struct {
	Id       int
	Name     string
	Birthday time.Time
}
