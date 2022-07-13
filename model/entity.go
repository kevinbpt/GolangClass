package model

import "time"

type User struct {
	Id        int       `json:"Id,omitempty"`
	Username  string    `json:"Username,omitempty"`
	Email     string    `json:"Email,omitempty"`
	Password  string    `json:"Password,omitempty"`
	Age       int       `json:"Age,omitempty"`
	CreatedAt time.Time `json:"CreatedAt,omitempty"`
	UpdatedAt time.Time `json:"UpdatedAt,omitempty"`
}

type Orders struct {
	OrderId      int       `json:"OrderId,omitempty"`
	CustomerName string    `json:"CustomerName,omitempty"`
	OrderedAt    time.Time `json:"OrderedAt,omitempty"`
	Item         []Items   `json:"Item,omitempty"`
}

type Items struct {
	ItemId      int
	ItemCode    string
	Description string
	Quantity    int
	OrderId     int
}
