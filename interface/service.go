package interfaces

import (
	"encoding/json"
	"fmt"
	"log"
	"projek-pertama/model"
)

type UserServiceIface interface {
	Register(user *model.User)
}

type UserSvc struct{}

func NewUserService() UserServiceIface {
	return &UserSvc{}
}

func (u *UserSvc) Register(user *model.User) {

	x, err := json.Marshal(user)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(x))
}
