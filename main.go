package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	interfaces "projek-pertama/interface"
	"projek-pertama/model"
	"strconv"
	"sync"
	"time"

	_ "github.com/denisenkom/go-mssqldb"
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

var wg sync.WaitGroup
var db *sql.DB

func dbConn() *sql.DB {
	var err error
	// connString := fmt.Sprintf("server=%s;port=%d; trusted_connection=yes/golang_db", server, port)
	// db, err = sql.Open("sqlserver", connString)
	db, err := sql.Open("sqlserver", "server=localhost;database=golang_db;trusted_connection=yes")
	if err != nil {
		log.Fatal("Error creating connection pool: " + err.Error())
	}
	if x := db.Ping(); x != nil {
		log.Fatal(x)
	}
	return db
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

	// var name = []string{"a", "b", "c", "d", "e", "f"}

	// for i := 0; i < len(name); i++ {
	// 	go print(name[i])
	// }
	// time.Sleep(1 * time.Second)

	// var err error
	// connString := fmt.Sprintf("server=%s;port=%d; trusted_connection=yes", server, port)
	// db, err = sql.Open("sqlserver", connString)
	// if err != nil {
	// 	log.Fatal("Error creating connection pool: " + err.Error())
	// }
	// if x := db.Ping(); x != nil {
	// 	log.Fatal(x)
	// }

	//defer db.Close()

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
	// for _, value := range users {
	// 	if value.Id == id {
	// 		json.NewEncoder(w).Encode(value)
	// 	}
	// }

	db := dbConn()
	var user model.User

	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()

	query := "SELECT id, username, email, password, age FROM dbo.MsUser WHERE Id = @Id"
	err := db.QueryRowContext(ctx, query, sql.Named("Id", id)).Scan(&user.Id, &user.Username, &user.Password, &user.Email, &user.Age)
	if err != nil {
		panic(err)
	}
	db.Close()
	json.NewEncoder(w).Encode(user)
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	users := []model.User{}
	query := "SELECT id, username, email, password, age FROM dbo.MsUser"
	searched, err := db.Query(query)
	if err != nil {
		panic(err)
	}

	for searched.Next() {
		var user model.User
		err := searched.Scan(&user.Id, &user.Username, &user.Password, &user.Email, &user.Age)
		if err != nil {
			panic(err)
		}
		users = append(users, user)
	}
	db.Close()
	json.NewEncoder(w).Encode(users)
}

func createUsers(w http.ResponseWriter, r *http.Request) {
	// var user = &model.User{}
	// if err := json.NewDecoder(r.Body).Decode(user); err != nil {
	// 	json.NewEncoder(w).Encode(err)
	// 	log.Fatal(err)
	// } else {
	// 	user.CreatedAt = time.Now()
	// 	users = append(users, user)
	// 	fmt.Fprint(w, "Success Create")
	// }
	var user = &model.User{}

	if err := json.NewDecoder(r.Body).Decode(user); err != nil {
		json.NewEncoder(w).Encode(err)
		log.Fatal(err)
	} else {
		db := dbConn()
		query := "INSERT INTO MsUser (username, password, email, age, createdate, updatedate) VALUES(@username, @password, @email, @age, @createdate, @updatedate)"
		ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancelfunc()

		res, err := db.ExecContext(ctx, query,
			sql.Named("username", user.Username),
			sql.Named("password", user.Password),
			sql.Named("email", user.Email),
			sql.Named("age", user.Age),
			sql.Named("createdate", time.Now()),
			sql.Named("updatedate", nil))
		if err != nil {
			log.Fatal(err)
		}
		res.LastInsertId()
		res.RowsAffected()
		defer db.Close()
		w.Write([]byte("User added successfully"))
	}

}

func updateUser(w http.ResponseWriter, r *http.Request, id int) {
	// for _, value := range users {
	// 	if value.Id == id {
	// 		var user = &model.User{}
	// 		if err := json.NewDecoder(r.Body).Decode(user); err != nil {
	// 			json.NewEncoder(w).Encode(err)
	// 			log.Fatal(err)
	// 		} else {
	// 			// value = &model.User{
	// 			// 	Id:        user.Id,
	// 			// 	Username:  user.Username,
	// 			// 	Email:     user.Email,
	// 			// 	Password:  user.Password,
	// 			// 	Age:       user.Age,
	// 			// 	UpdatedAt: time.Now(),
	// 			// }
	// 			value.Id = user.Id
	// 			value.Username = user.Username
	// 			value.Email = user.Email
	// 			value.Password = user.Password
	// 			value.Age = user.Age
	// 			value.UpdatedAt = time.Now()
	// 			fmt.Fprint(w, "Success Edit")
	// 		}
	// 	}
	// }
	//fmt.Println(users)

	var user = &model.User{}

	if err := json.NewDecoder(r.Body).Decode(user); err != nil {
		json.NewEncoder(w).Encode(err)
		log.Fatal(err)
	} else {
		db := dbConn()
		query := "UPDATE dbo.MsUser SET username = @username, email = @email, password = @password, age = @age, updatedate = @updatedate WHERE id = @id"
		ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancelfunc()

		res, err := db.ExecContext(ctx, query,
			sql.Named("username", user.Username),
			sql.Named("password", user.Password),
			sql.Named("email", user.Email),
			sql.Named("age", user.Age),
			sql.Named("updatedate", time.Now()),
			sql.Named("id", id))
		if err != nil {
			log.Fatal(err)
		}
		res.LastInsertId()
		res.RowsAffected()
		defer db.Close()
		w.Write([]byte("User updated successfully"))
	}
}

func deleteUser(w http.ResponseWriter, r *http.Request, id int) {
	// for i, value := range users {
	// 	if value.Id == id {
	// 		users = users[:i+copy(users[i:], users[i+1:])]
	// 		fmt.Fprint(w, "Success Delete")
	// 	}
	// }
	db := dbConn()
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()

	_, err := db.ExecContext(ctx, "DELETE FROM dbo.MsUser WHERE id=@id",
		sql.Named("id", id))

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	w.Write([]byte("User deleted successfully"))
}
