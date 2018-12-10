package main

import (
	"log"
	"net/http"

	mux "github.com/gorilla/mux"
	"github.com/iecheniq/go_bootcamp/cartutils"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	if err := cartutils.OpenCartDB(); err != nil {
		log.Fatal(err)
	}
	defer cartutils.CloseCartDB()
	setRouts()
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func setRouts() {
	r := mux.NewRouter()
	r.HandleFunc("/", cartutils.HomeHandler).
		Methods("GET").
		Name("home")
	r.HandleFunc("/carts", cartutils.CartsHandler).
		Methods("GET", "POST").
		Name("carts")
	cartSubrouter := r.PathPrefix("/carts/{cartID:[0-9]+}").Subrouter()
	cartSubrouter.HandleFunc("/", cartutils.CartHandler).
		Methods("GET", "DELETE").
		Name("cart_details")
	cartSubrouter.HandleFunc("/items", cartutils.CartItemsHandler).
		Methods("GET", "POST", "DELETE").
		Name("items")
	cartSubrouter.HandleFunc("/items/{itemID:[0-9]+}", cartutils.CartItemHandler).
		Methods("GET", "DELETE", "PATCH").
		Name("item_details")
	http.Handle("/", r)
}
