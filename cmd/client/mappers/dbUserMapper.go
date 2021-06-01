package mappers

import (
	"github.com/google/uuid"
	. "userService/pkg/models/dbModels"
	. "userService/pkg/models/requestModels"
)

func DbUserMap(userRequest CreateUserRequest, id uuid.UUID) User {
	return User{
		Id:       id,
		Name:     userRequest.Name,
		Birthday: userRequest.Birthday,
	}
}
