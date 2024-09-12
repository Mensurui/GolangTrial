package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	db := database{"shoes": 50, "shirt": 40}
	mux := http.NewServeMux()
	mux.Handle("/list", http.HandlerFunc(db.list))
	mux.Handle("/price", http.HandlerFunc(db.price))
	mux.Handle("/add", http.HandlerFunc(db.addItem))
	log.Fatal(http.ListenAndServe("localhost:8000", mux))
}

type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

type database map[string]dollars

func (db database) addItem(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	priceStr := req.URL.Query().Get("price")

	var price dollars
	_, err := fmt.Scanf(priceStr, "%f", &price)

	if item == "" || err != nil {
		w.WriteHeader(http.StatusBadRequest) // 400
		fmt.Fprintf(w, "Invalid item or price.\n")
		return
	}
	db[item] = price

	fmt.Fprintf(w, "Added item: %s with price: %s\n", item, price)
}

func (db database) list(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

func (db database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}
	fmt.Fprintf(w, "%s\n", price)
}
