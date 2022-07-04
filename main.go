package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	interfaces "projek-pertama/interface"
	"projek-pertama/model"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

var PORT = ":8080"

var users = []*model.User{
	{
		Id:       1,
		Username: "andi123",
		Email:    "andi123@gmail.com",
		Password: "password123",
		Age:      9,
	},
	{
		Id:       2,
		Username: "budi123",
		Email:    "budi123@gmail.com",
		Password: "password123",
		Age:      9,
	},
	{
		Id:       3,
		Username: "cantya123",
		Email:    "cantya123@gmail.com",
		Password: "password123",
		Age:      9,
	},
	{
		Id:       4,
		Username: "cantya123",
		Email:    "dantya123@gmail.com",
		Password: "password123",
		Age:      9,
	},
	{
		Id:       5,
		Username: "2312",
		Email:    "4415@gmail.com",
		Password: "password123",
		Age:      9,
	},
}

func main() {
	// http.HandleFunc("/", greet)
	// http.HandleFunc("/register", reg)
	// http.ListenAndServe(PORT, nil)

	r := mux.NewRouter()
	r.HandleFunc("/users", UsersHandler)
	r.HandleFunc("/users/{id}", UsersHandler)
	http.Handle("/", r)
	http.ListenAndServe(PORT, nil)
}

func greet(w http.ResponseWriter, r *http.Request) {
	msg := "Hello world"
	w.Header().Add("asd", "aq12")
	fmt.Fprint(w, msg)
}

func reg(w http.ResponseWriter, r *http.Request) {
	// userSvc := interfaces.NewUserService()
	// now := time.Now()
	// temp := userSvc.Register(&model.User{
	// 	Id:        1,
	// 	Username:  "Kevin",
	// 	Email:     "Kevin@test.com",
	// 	Password:  "123123",
	// 	Age:       22,
	// 	CreatedAt: now,
	// 	UpdatedAt: now,
	// })

	// w.Write(temp)

	//------------------------------------------------------------------

	decoder := json.NewDecoder(r.Body)
	var user model.User
	if err := decoder.Decode(&user); err != nil {
		//w.Write([]byte("error"))
		fmt.Fprint(w, "err")
		return
	}

	userSvc := interfaces.NewUserService()
	temp, err := userSvc.Register(&user)

	w.Header().Add("Content-Type", "application/json")
	if err != nil {
		w.Write([]byte(err.Error()))
	} else {
		w.Write(temp)
	}

}

func UsersHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	var tempId, _ = strconv.Atoi(id)

	switch r.Method {
	case http.MethodGet:
		if id != "" { // get by id
			getUsersById(w, r, tempId)
		} else { // get all
			getUsers(w, r)
		}
	case http.MethodPost:
		createUsers(w, r)
	case http.MethodPut:
		updateUser(w, r, tempId)
	case http.MethodDelete:
		deleteUser(w, r, tempId)
	}
}

func getUsersById(w http.ResponseWriter, r *http.Request, id int) {
	for _, value := range users {
		if value.Id == id {
			json.NewEncoder(w).Encode(value)
		}
	}
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(users)
}

func createUsers(w http.ResponseWriter, r *http.Request) {
	//var user *model.User
	var user = &model.User{}
	if err := json.NewDecoder(r.Body).Decode(user); err != nil {
		json.NewEncoder(w).Encode(err)
		log.Fatal(err)
	} else {
		user.CreatedAt = time.Now()
		users = append(users, user)
		fmt.Fprint(w, "Success Create")
	}
}

func updateUser(w http.ResponseWriter, r *http.Request, id int) {
	for _, value := range users {
		if value.Id == id {
			var user = &model.User{}
			if err := json.NewDecoder(r.Body).Decode(user); err != nil {
				json.NewEncoder(w).Encode(err)
				log.Fatal(err)
			} else {
				// value = &model.User{
				// 	Id:        user.Id,
				// 	Username:  user.Username,
				// 	Email:     user.Email,
				// 	Password:  user.Password,
				// 	Age:       user.Age,
				// 	UpdatedAt: time.Now(),
				// }
				value.Id = user.Id
				value.Username = user.Username
				value.Email = user.Email
				value.Password = user.Password
				value.Age = user.Age
				value.UpdatedAt = time.Now()
				fmt.Fprint(w, "Success Edit")
			}
		}
	}
	fmt.Println(users)
}

func deleteUser(w http.ResponseWriter, r *http.Request, id int) {
	for i, value := range users {
		if value.Id == id {
			users = users[:i+copy(users[i:], users[i+1:])]
			fmt.Fprint(w, "Success Delete")
		}
	}
}
