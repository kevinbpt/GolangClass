package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"projek-pertama/model"
	"strconv"
	"sync"
	"time"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/gorilla/mux"
)

var PORT = ":8088"

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
	db, err := sql.Open("sqlserver", "server=localhost;database=golang_db;trusted_connection=yes")
	if err != nil {
		log.Fatal("Error creating connection pool: " + err.Error())
	}
	// if x := db.Ping(); x != nil {
	// 	log.Fatal(x)
	// }
	return db
}

func main() {
	// http.HandleFunc("/", greet)
	// http.HandleFunc("/register", reg)
	// http.ListenAndServe(PORT, nil)

	db = dbConn()
	defer db.Close()

	r := mux.NewRouter()
	r.HandleFunc("/greet", greet)
	r.HandleFunc("/orders", UsersHandler)
	r.HandleFunc("/orders/{id}", UsersHandler)
	http.Handle("/", r)
	http.ListenAndServe(PORT, nil)

}

func greet(w http.ResponseWriter, r *http.Request) {
	msg := "Hello world"
	w.Header().Add("asd", "aq12")
	fmt.Fprint(w, msg)
}

func UsersHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	var tempId, _ = strconv.Atoi(id)

	switch r.Method {
	case http.MethodGet:
		if id != "" { // get by id
			getOrdersById(w, r, tempId)
		} else { // get all
			getOrders(w, r)
		}
	case http.MethodPost:
		createOrder(w, r)
	case http.MethodPut:
		updateOrder(w, r, tempId)
	case http.MethodDelete:
		deleteOrder(w, r, tempId)
	}
}

func getOrdersById(w http.ResponseWriter, r *http.Request, id int) {

	var order model.Orders
	var items []model.Items

	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()

	query := "SELECT order_id, customer_name, ordered_at FROM dbo.orders WHERE order_id = @Id"
	err := db.QueryRowContext(ctx, query, sql.Named("Id", id)).Scan(&order.OrderId, &order.CustomerName, &order.OrderedAt)
	if err != nil {
		panic(err)
	}

	queryItem := "SELECT item_id, item_code, description, quantity, order_id FROM dbo.items WHERE order_id = @Id"
	searchedItem, err2 := db.Query(queryItem, sql.Named("Id", order.OrderId))
	if err2 != nil {
		panic(err2)
	}

	for searchedItem.Next() {
		var item model.Items
		err3 := searchedItem.Scan(&item.ItemId, &item.ItemCode, &item.Description, &item.Quantity, &item.OrderId)
		if err3 != nil {
			panic(err3)
		}
		items = append(items, item)
	}
	order.Item = items

	json.NewEncoder(w).Encode(order)
}

func getOrders(w http.ResponseWriter, r *http.Request) {
	orders := []model.Orders{}
	query := "SELECT order_id, customer_name, ordered_at FROM dbo.orders"
	searched, err := db.Query(query)
	if err != nil {
		panic(err)
	}

	for searched.Next() {
		var order model.Orders
		var items []model.Items

		err := searched.Scan(&order.OrderId, &order.CustomerName, &order.OrderedAt)
		if err != nil {
			panic(err)
		}
		query := "SELECT item_id, item_code, description, quantity, order_id FROM dbo.items WHERE order_id = @Id"
		searchedItem, err2 := db.Query(query, sql.Named("Id", order.OrderId))
		if err2 != nil {
			panic(err2)
		}

		for searchedItem.Next() {
			var item model.Items
			err3 := searchedItem.Scan(&item.ItemId, &item.ItemCode, &item.Description, &item.Quantity, &item.OrderId)
			if err3 != nil {
				panic(err3)
			}
			items = append(items, item)
		}
		order.Item = items
		orders = append(orders, order)
	}
	json.NewEncoder(w).Encode(orders)
}

func createOrder(w http.ResponseWriter, r *http.Request) {
	var order = &model.Orders{}
	if err := json.NewDecoder(r.Body).Decode(order); err != nil {
		json.NewEncoder(w).Encode(err)
		log.Fatal(err)
	} else {
		query := "INSERT INTO dbo.orders (customer_name, ordered_at) VALUES(@customer_name, @ordered_at); select order_id = convert(bigint, SCOPE_IDENTITY())"
		queryItem := "INSERT INTO dbo.items (item_code, description, quantity, order_id) VALUES(@item_code, @description, @quantity, @order_id)"
		orderedAt := time.Now()
		ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancelfunc()
		res, err := db.QueryContext(ctx, query,
			sql.Named("customer_name", order.CustomerName),
			sql.Named("ordered_at", orderedAt))
		if err != nil {
			log.Fatal(err)
		}

		var orderId int
		for res.Next() {
			err := res.Scan(&orderId)
			if err != nil {
				log.Fatal(err)
			}
		}

		for _, item := range order.Item {
			_, err := db.ExecContext(ctx, queryItem,
				sql.Named("item_code", item.ItemCode),
				sql.Named("description", item.Description),
				sql.Named("quantity", item.Quantity),
				sql.Named("order_id", orderId))
			if err != nil {
				log.Fatal(err)
			}
		}
		w.Write([]byte("Order added successfully"))
	}

}

func updateOrder(w http.ResponseWriter, r *http.Request, id int) {

	var order = &model.Orders{}
	if err := json.NewDecoder(r.Body).Decode(order); err != nil {
		json.NewEncoder(w).Encode(err)
		log.Fatal(err)
	} else {
		query := "UPDATE dbo.orders SET customer_name = @customer_name, ordered_at = @ordered_at WHERE order_id = @id"
		query2 := "UPDATE dbo.items SET item_code = @item_code, description = @description, quantity = @quantity WHERE item_id = @item_id"
		ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancelfunc()

		_, err := db.ExecContext(ctx, query,
			sql.Named("customer_name", order.CustomerName),
			sql.Named("ordered_at", order.OrderedAt),
			sql.Named("id", id))
		if err != nil {
			log.Fatal(err)
		}

		for _, item := range order.Item {
			_, err := db.ExecContext(ctx, query2,
				sql.Named("item_code", item.ItemCode),
				sql.Named("description", item.Description),
				sql.Named("quantity", item.Quantity),
				sql.Named("item_id", item.ItemId))
			if err != nil {
				log.Fatal(err)
			}
		}
		w.Write([]byte("Order updated successfully"))
	}
}

func deleteOrder(w http.ResponseWriter, r *http.Request, id int) {

	ctx, cancelfunc := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancelfunc()

	_, err := db.ExecContext(ctx, "DELETE FROM dbo.orders WHERE order_id=@id",
		sql.Named("id", id))

	if err != nil {
		log.Fatal(err)
	}

	_, err2 := db.ExecContext(ctx, "DELETE FROM dbo.items WHERE order_id=@id",
		sql.Named("id", id))

	if err2 != nil {
		log.Fatal(err2)
	}
	w.Write([]byte("Order deleted successfully"))
}
