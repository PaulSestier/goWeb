package handlers

// import (
// 	"bootcamp/goweb/internal"
// 	"encoding/json"
// 	"fmt"
// 	"io"
// 	"net/http"
// 	"os"
// 	"strconv"

// 	"github.com/go-chi/chi/v5"
// )

// type ProductsSlice struct {
// 	Products []internal.Product
// }

// // function to create a new slice of products
// func NewProductSlice() *ProductsSlice {
// 	return &ProductsSlice{
// 		Products: make([]internal.Product, 0),
// 	}
// }

// // function to load products slice from json file
// func (p *ProductsSlice) LoadProducts(filename string) error {
// 	//open json file
// 	file, err := os.Open(filename)
// 	if err != nil {
// 		return err
// 	}
// 	defer file.Close()

// 	//read file
// 	byteFile, err := io.ReadAll(file)
// 	if err != nil {
// 		fmt.Println(err)
// 		return err
// 	}

// 	//unmarshal file into slice of model Product
// 	err = json.Unmarshal(byteFile, &p.Products)

// 	if err != nil {
// 		fmt.Println(err)
// 		return err
// 	}
// 	fmt.Println("Product slice loaded successfully")
// 	return nil
// }

// // function to return a ping response
// func (p *ProductsSlice) Ping() http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		w.Header().Add("Content-Type", "text/plain")
// 		w.WriteHeader(http.StatusOK)
// 		w.Write([]byte("pong"))
// 	}
// }

// // function to return a list of products
// func (p *ProductsSlice) GetProducts() http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		w.Header().Add("Content-Type", "application/json")
// 		w.WriteHeader(http.StatusOK)
// 		json.NewEncoder(w).Encode(p.Products)
// 	}
// }

// // function to return a product by id
// func (p *ProductsSlice) GetProductById() http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		id := chi.URLParam(r, "productId")
// 		idvalue, err := strconv.Atoi(id)
// 		if err != nil {
// 			w.Header().Add("Content-Type", "text/plain")
// 			w.WriteHeader(http.StatusBadRequest)
// 			w.Write([]byte("invalid id"))
// 			return
// 		}

// 		for _, product := range p.Products {
// 			if product.Id == idvalue {
// 				w.Header().Add("Content-Type", "application/json")
// 				w.WriteHeader(http.StatusOK)
// 				json.NewEncoder(w).Encode(product)
// 			}
// 		}

// 	}
// }

// // function that returns products with higher price than the one passed as parameter
// func (p *ProductsSlice) PriceHigherThan() http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		productsGt := NewProductSlice()

// 		price := r.URL.Query().Get("priceGt")
// 		pricevalue, err := strconv.ParseFloat(price, 64)
// 		if err != nil {
// 			w.Header().Add("Content-Type", "text/plain")
// 			w.WriteHeader(http.StatusBadRequest)
// 			w.Write([]byte("invalid price"))
// 			return
// 		}

// 		for _, product := range p.Products {
// 			if product.Price > pricevalue {
// 				productsGt.Products = append(productsGt.Products, product)
// 			}
// 		}
// 		w.Header().Add("Content-Type", "application/json")
// 		w.WriteHeader(http.StatusOK)
// 		json.NewEncoder(w).Encode(productsGt.Products)
// 	}
// }

// // function to create new product
// func (p *ProductsSlice) CreateProduct() http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		var product internal.Product
// 		err := json.NewDecoder(r.Body).Decode(&product)
// 		if err != nil {
// 			w.Header().Add("Content-Type", "text/plain")
// 			w.WriteHeader(http.StatusBadRequest)
// 			w.Write([]byte("invalid product"))
// 			return
// 		}
// 		err = product.ValidateFields()
// 		if err != nil {
// 			w.Header().Add("Content-Type", "text/plain")
// 			w.WriteHeader(http.StatusBadRequest)
// 			w.Write([]byte(err.Error()))
// 			return
// 		}

// 		p.Products = append(p.Products, product)
// 		w.Header().Add("Content-Type", "application/json")
// 		w.WriteHeader(http.StatusCreated)
// 		json.NewEncoder(w).Encode(product)
// 	}
// }
