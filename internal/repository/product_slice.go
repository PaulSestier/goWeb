package repository

import (
	"bootcamp/goweb/internal"
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type ProductsSlice struct {
	Products []internal.Product
	lastId   int
}

// function to create a new slice of products
func NewProductSlice() *ProductsSlice {
	return &ProductsSlice{
		Products: make([]internal.Product, 0),
	}
}

func (p *ProductsSlice) LoadProducts(filename string, lastId int) (err error) {
	//open json file
	file, err := os.Open(filename)
	if err != nil {
		return
	}
	defer file.Close()

	//read file
	byteFile, err := io.ReadAll(file)
	if err != nil {
		fmt.Println(err)
		return
	}

	//unmarshal file into slice of model Product
	err = json.Unmarshal(byteFile, &p.Products)
	p.lastId = p.Products[len(p.Products)-1].Id

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Product slice loaded successfully, lastId:", p.lastId)
	return
}

// function that returns all products
func (p *ProductsSlice) GetProducts() ([]internal.Product, error) {
	return p.Products, nil
}

func (p *ProductsSlice) GetProductById(id int) (*internal.Product, error) {
	for _, product := range p.Products {
		if product.Id == id {
			return &product, nil
		}
	}
	return nil, internal.ErrProductNotFound
}

func (p *ProductsSlice) CreateProduct(product *internal.Product) error {
	p.lastId++
	product.Id = p.lastId
	p.Products = append(p.Products, *product)
	return nil
}
