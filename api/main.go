package main

import (
	"bootcamp/goweb/model"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
)

func main() {
	//call function to load slice with json file
	var products []model.Product
	err := loadSlice(&products)
	if err != nil {
		fmt.Println(err)
		return
	}

	handlerPing := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "text/plain")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("pong"))
	}

	handlerProducts := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(products)
	}

	// router
	router := chi.NewRouter()
	router.Get("/ping", handlerPing)
	router.Get("/products", handlerProducts)

	if err := http.ListenAndServe(":8080", router); err != nil {
		println("Error starting server: ", err.Error())
		return
	}

}

func loadSlice(products *[]model.Product) error {
	//open json file
	file, err := os.Open("products.json")
	if err != nil {
		return err
	}
	defer file.Close()

	//read file
	byteFile, err := io.ReadAll(file)
	if err != nil {
		fmt.Println(err)
		return err
	}

	//unmarshal file into slice of model Product
	err = json.Unmarshal(byteFile, &products)

	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println("Product slice loaded successfully")
	return nil
}
