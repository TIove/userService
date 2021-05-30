package requestModels

import "time"

type CreateUserRequest struct {
	Name string
	Birthday time.Time
}