package interfaces

import (
	"encoding/json"
	"errors"
	"log"
	"projek-pertama/model"
)

type UserServiceIface interface {
	Register(user *model.User) ([]byte, error)
}

type UserSvc struct{}

func NewUserService() UserServiceIface {
	return &UserSvc{}
}

func (u *UserSvc) Register(user *model.User) ([]byte, error) {

	x, err := json.Marshal(user)
	if err != nil {
		log.Fatal(err)
	}
	if len(user.Username) < 1 {
		return nil, errors.New("Input username")
	}
	return x, nil
}
