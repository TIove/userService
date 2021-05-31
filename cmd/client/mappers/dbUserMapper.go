package mappers

import (
	. "userService/pkg/models/dbModels"
	. "userService/pkg/models/requestModels"
)

func DbUserMap(userRequest CreateUserRequest, id int) User {
	return User{
		Id:       id,
		Name:     userRequest.Name,
		Birthday: userRequest.Birthday,
	}
}
