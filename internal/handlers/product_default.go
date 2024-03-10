package handlers

import (
	"bootcamp/goweb/internal"
	"net/http"
	"strconv"

	"github.com/bootcamp-go/web/response"
	"github.com/go-chi/chi/v5"
)

type DefaulProduct struct {
	sv internal.ProductService
}

func NewDefaultProduct(sv internal.ProductService) *DefaulProduct {
	return &DefaulProduct{
		sv: sv,
	}
}

func (p *DefaulProduct) GetProducts() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		products, err := p.sv.GetProducts()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		response.JSON(w, http.StatusOK, products)
	}
}

func (p *DefaulProduct) GetProductById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "productId")
		idValue, err := strconv.Atoi(id)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(internal.ErrorProductInvalidId.Error()))
			return
		}
		product, err := p.sv.GetProductById(idValue)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		response.JSON(w, http.StatusOK, product)
	}
}
