package main

import (
	"bootcamp/goweb/internal/handlers"
	"bootcamp/goweb/internal/repository"
	"bootcamp/goweb/internal/service"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	//create new slice of products
	rp := repository.NewProductSlice()
	rp.LoadProducts("./docs/db/products.json", 0)
	sv := service.NewProductDefault(rp)
	hd := handlers.NewDefaultProduct(sv)

	// router
	router := chi.NewRouter()
	router.Route("/products", func(r chi.Router) {
		r.Get("/", hd.GetProducts())
		r.Get("/{productId}", hd.GetProductById())
	})
	// router.Get("/ping", store.Ping())
	// router.Get("/products", store.GetProducts())
	// router.Get("/products/{productId}", store.GetProductById())
	// router.Get("/products/search", store.PriceHigherThan())
	// router.Post("/products", store.CreateProduct())

	if err := http.ListenAndServe(":8080", router); err != nil {
		println("Error starting server: ", err.Error())
		return
	}

}
