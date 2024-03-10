package internal

import "errors"

type Product struct {
	Id          int
	Name        string
	Quantity    int
	CodeValue   string
	IsPublished bool
	Expiration  string
	Price       float64
}

var (
	ErrProductNotFound     = errors.New("product not found")
	ErrProductDuplicated   = errors.New("product duplicated")
	ErrProductInvalidField = errors.New("product invalid field")
	ErrorProductInvalidId  = errors.New("product invalid id")
)

// ProductRepository is an interface that defines the methods that a product repository must implement
type ProductRepository interface {
	GetProductById(id int) (*Product, error)
	GetProducts() ([]Product, error)
	// GetProductsByPrice(price float64) ([]Product, error)
	CreateProduct(product *Product) error
}

// TaskService is an interface that defines the methods that a task service must implement
type ProductService interface {
	GetProductById(id int) (*Product, error)
	GetProducts() ([]Product, error)
	// GetProductsByPrice(price float64) ([]Product, error)
	CreateProduct(product *Product) error
}

//"id":1,"name":"Oil - Margarine","quantity":439,"code_value":"S82254D","is_published":true,"expiration":"15/12/2021","price":71.42

// func (p *Product) ValidateFields() error {
// 	if p.Id == 0 {
// 		return errors.New("invalid id")
// 	}
// 	if p.Name == "" {
// 		return errors.New("invalid name")
// 	}
// 	if p.Quantity == 0 {
// 		return errors.New("invalid quantity")
// 	}
// 	if p.CodeValue == "" {
// 		return errors.New("invalid code_value")
// 	}
// 	if p.Expiration == "" {
// 		return errors.New("invalid expiration")
// 	}
// 	if p.Price == 0 {
// 		return errors.New("invalid price")
// 	}
// 	return nil
// }
