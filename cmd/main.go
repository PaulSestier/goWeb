package main

import (
	"bootcamp/goweb/internal/handlers"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	//create new slice of products
	store := handlers.NewProductSlice()

	//load products from json file
	err := store.LoadProducts("products.json")
	if err != nil {
		fmt.Println(err)
		return
	}

	// router
	router := chi.NewRouter()
	router.Get("/ping", store.Ping())
	router.Get("/products", store.GetProducts())
	router.Get("/products/{productId}", store.GetProductById())

	if err := http.ListenAndServe(":8080", router); err != nil {
		println("Error starting server: ", err.Error())
		return
	}

}
