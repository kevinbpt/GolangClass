package main

import (
	interfaces "projek-pertama/interface"
	"projek-pertama/model"
	"time"
)

func main() {
	userSvc := interfaces.NewUserService()
	now := time.Now()
	userSvc.Register(&model.User{
		Id:        1,
		Username:  "Kevin",
		Email:     "Kevin@test.com",
		Password:  "123123",
		Age:       22,
		CreatedAt: now,
		UpdatedAt: now,
	})
}
