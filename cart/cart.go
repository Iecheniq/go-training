package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	mux "github.com/gorilla/mux"
)

import "database/sql"
import _ "github.com/go-sql-driver/mysql"

var db *sql.DB

//Item is the model used for articles in a cart
type Item struct {
	Id     int64
	Title  string
	Price  float64
	CartID int64
}

//Cart is the model that describes a shopping cart
type Cart struct {
	Id    int64
	Items []Item
}

func main() {

	db, _ = sql.Open("mysql", "iecheniq:HoUsE22$@tcp(localhost:3306)/shopping_cart")
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	setRouts()
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func setRouts() {
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler).
		Methods("GET").
		Name("home")
	r.HandleFunc("/carts", CartsHandler).
		Methods("GET", "POST").
		Name("carts")
	cartSubrouter := r.PathPrefix("/carts/{cartID:[0-9]+}").Subrouter()
	cartSubrouter.HandleFunc("/", CartHandler).
		Methods("GET", "DELETE").
		Name("cart_details")
	cartSubrouter.HandleFunc("/items", CartItemsHandler).
		Methods("GET", "POST", "DELETE").
		Name("items")
	cartSubrouter.HandleFunc("/items/{itemID:[0-9]+}", CartItemHandler).
		Methods("GET", "DELETE", "UPDATE", "PATCH").
		Name("item_details")
	http.Handle("/", r)
}

//HomeHandler handles the root URL
//GET: Description of the shopping cart API
func HomeHandler(w http.ResponseWriter, r *http.Request) {

	if _, err := w.Write([]byte(fmt.Sprintf("Welcome to the shopping cart API"))); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

//CartsHandler handles all existing carts
//GET: List all carts
//POST: Create a new cart
func CartsHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {

		carts := make(map[string]Cart)
		rows, err := db.Query("SELECT carts.*, items.* FROM carts LEFT JOIN items ON carts.id = items.cart_id")
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()
		var cartID sql.NullInt64
		var itemID sql.NullInt64
		var itemTitle sql.NullString
		var itemPrice sql.NullFloat64
		var itemCart sql.NullInt64

		for rows.Next() {
			err := rows.Scan(&cartID, &itemID, &itemTitle, &itemPrice, &itemCart)
			if err != nil {
				log.Fatal(err)
			}
			cart := Cart{
				Id:    cartID.Int64,
				Items: make([]Item, 0),
			}
			if itemCart.Valid {
				item := Item{
					Id:     itemID.Int64,
					Title:  itemTitle.String,
					Price:  itemPrice.Float64,
					CartID: itemCart.Int64,
				}
				if c, ok := carts["Cart "+strconv.Itoa(int(cart.Id))]; ok {
					c.Items = append(c.Items, item)
					carts["Cart "+strconv.Itoa(int(cart.Id))] = c
				} else {
					cart.Items = append(cart.Items, item)
					carts["Cart "+strconv.Itoa(int(cart.Id))] = cart
				}
			} else {
				carts["Cart "+strconv.Itoa(int(cart.Id))] = cart
			}
		}
		err = rows.Err()
		if err != nil {
			log.Fatal(err)
		}
		Jcarts, err := json.Marshal(carts)
		if err != nil {
			panic(err)
		}
		w.WriteHeader(http.StatusOK)
		if _, err := w.Write(Jcarts); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	} else if r.Method == "POST" {
		res, err := db.Exec("INSERT INTO carts() VALUES()")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Fatal(err)
		}

		lastId, err := res.LastInsertId()
		if err != nil {
			log.Fatal(err)
		}
		rowCnt, err := res.RowsAffected()
		if err != nil {
			log.Fatal(err)
		}
		w.WriteHeader(http.StatusCreated)
		if _, err := w.Write([]byte(fmt.Sprintf("Cart with ID %v has been created", lastId))); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		log.Printf("ID = %d, affected = %d\n", lastId, rowCnt)
	} else {
		http.Error(w, "Allowed methods: GET, POST", http.StatusMethodNotAllowed)
	}
}

//CartHandler handles a cart.
//GET: Get a specific cart
//DELETE: Delete the cart
func CartHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	if r.Method == "GET" {
		cart := Cart{
			Id:    0,
			Items: make([]Item, 0),
		}
		item := Item{
			Id:     0,
			Title:  "",
			Price:  0.0,
			CartID: 0,
		}
		err := db.QueryRow("SELECT * FROM carts WHERE id = ? ", vars["cartID"]).Scan(&cart.Id)
		if err != nil {
			log.Fatal(err)
		}
		stmt, err := db.Prepare("SELECT * FROM items WHERE cart_id = ?")
		if err != nil {
			log.Fatal(err)
		}
		defer stmt.Close()
		rows, err := stmt.Query(vars["cartID"])
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()
		for rows.Next() {
			err := rows.Scan(&item.Id, &item.Title, &item.Price, &item.CartID)
			if err != nil {
				log.Fatal(err)
			}
			cart.Items = append(cart.Items, item)
		}
		err = rows.Err()
		if err != nil {
			log.Fatal(err)
		}
		Jcarts, err := json.Marshal(cart)
		if err != nil {
			panic(err)
		}
		w.WriteHeader(http.StatusOK)
		if _, err := w.Write(Jcarts); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	} else if r.Method == "DELETE" {
		res, err := db.Exec("DELETE FROM carts WHERE id = ? ", vars["cartID"])
		if err != nil {
			log.Fatal(err)
		}
		w.WriteHeader(http.StatusOK)
		if _, err := w.Write([]byte(fmt.Sprintf("Cart with ID %v has been deleted", vars["cartID"]))); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		rowCnt, err := res.RowsAffected()
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("%d rows affected\n", rowCnt)
	} else {
		http.Error(w, "Allowed methods: GET, DELETE", http.StatusMethodNotAllowed)
	}

}

//CartItemsHandler handles the items of a cart.
//GET: Get all items of a cart
//POST: Add a new item to the cart
//DELETE: Delete all items of a cart
func CartItemsHandler(w http.ResponseWriter, r *http.Request) {

}

//CartItemHandler handles a specific item in a cart
//GET: Get a specific item
//DELETE: Delete item
//UPDATE: Update item
//PATCH: Modify item
func CartItemHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	if r.Method == "GET" {
		response, _ := http.Get(fmt.Sprintf("http://challenge.getsandbox.com/articles/%s", vars["itemID"]))
		body, err := ioutil.ReadAll(response.Body)

		if err != nil {
			panic(err.Error())
		}
		if _, err := w.Write(body); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}
