package user

import (
	api "github.com/balticoldschool/user-service/.generated"
	rest_error "github.com/balticoldschool/user-service/src/utils/rest-error"
	"github.com/google/uuid"
)

type UserService interface {
	FindUserById(userId uuid.UUID) (*api.UserReadDto, *rest_error.RestError)
}

type service struct{}

func NewUserService() UserService {
	return service{}
}

func (s service) FindUserById(userId uuid.UUID) (*api.UserReadDto, *rest_error.RestError) {
	member := true

	user := api.UserReadDto{
		Id:        userId,
		FirstName: "Max",
		LastName:  "Mustermann",
		Email:     "max@mustermann.de",
		Member:    &member,
	}
	if userId.String() == "a1b2c3d4-e5f6-7890-1234-567890abcdef" {
		return &user, nil
	}

	return nil, rest_error.NewNotFoundError("User not found")
}
