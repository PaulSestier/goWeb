package service

import "bootcamp/goweb/internal"

type ProductsDefault struct {
	rp internal.ProductRepository
}

func NewProductDefault(rp internal.ProductRepository) *ProductsDefault {
	return &ProductsDefault{
		rp: rp,
	}
}

func (p *ProductsDefault) GetProducts() (products []internal.Product, err error) {
	products, err = p.rp.GetProducts()
	if err != nil {
		return nil, err
	}
	return
}

func (p *ProductsDefault) GetProductById(id int) (product *internal.Product, err error) {
	product, err = p.rp.GetProductById(id)
	if err != nil {
		return nil, err
	}

	return
}

func (p *ProductsDefault) CreateProduct(product *internal.Product) (err error) {
	err = p.rp.CreateProduct(product)
	return
}
