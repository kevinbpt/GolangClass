package model

import (
	"time"
)

type User struct {
	Id        int       `json:"Id,omitempty"`
	Username  string    `json:"Username,omitempty"`
	Password  string    `json:"Password,omitempty"`
	Email     string    `json:"Email,omitempty"`
	Age       int       `json:"Age,omitempty"`
	CreatedAt time.Time `json:"CreatedAt,omitempty"`
	UpdatedAt time.Time `json:"UpdatedAt,omitempty"`
}

type Auth struct {
	Username string
	Password string
}

type Token struct {
	Username string
	Token    string
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

type UserData struct {
	Id         int
	Uid        string
	First_name string
	Last_name  string
	Username   string
	Address    Address
}

type Address struct {
	City           string
	Street_name    string
	Street_address string
	Zip_code       string
	State          string
	Country        string
	Coordinates    Coordinates
}

type Coordinates struct {
	Lat float64
	Lng float64
}
