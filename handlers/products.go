package handlers

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/Omar-Belghaouti/microservices/data"
	"github.com/gorilla/mux"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) GetProducts(w http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle GET Products")
	lp := data.GetProducts()
	err := lp.ToJSON(w)
	if err != nil {
		http.Error(w, "Unable to marshal json", http.StatusInternalServerError)
		return
	}
}

func (p *Products) AddProduct(w http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POST Product")

	prod := r.Context().Value(ProductKey{}).(*data.Product)

	data.AddProduct(prod)
}

func (p *Products) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle PUT Product")

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Unable to convert id", http.StatusBadRequest)
	}

	prod := r.Context().Value(ProductKey{}).(*data.Product)

	err = data.UpdateProduct(id, prod)
	if err != nil {
		if err == data.ErrProductNotFound {
			http.Error(w, "Product not found", http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func (p *Products) getProductFromBody(w http.ResponseWriter, r *http.Request) (*data.Product, error) {
	defer r.Body.Close()

	prod := &data.Product{}
	err := prod.FromJSON(r.Body)
	if err != nil {
		return nil, err
	}
	return prod, nil
}

type ProductKey struct{}

func (p *Products) MiddlewareProductValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		prod, err := p.getProductFromBody(w, r)
		if err != nil {
			p.l.Println("[Error] deserializing product", err)
			http.Error(w, "Unable to unmarshal product json", http.StatusBadRequest)
			return
		}

		// validate product
		err = prod.Validate()
		if err != nil {
			p.l.Println("[Error] validating product", err)
			http.Error(w, "Unable to validate product", http.StatusBadRequest)
			return
		}

		ctx := context.WithValue(r.Context(), ProductKey{}, prod)
		req := r.WithContext(ctx)
		next.ServeHTTP(w, req)
	})
}
