package dbModels

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	Id       uuid.UUID
	Name     string
	Birthday time.Time
}
