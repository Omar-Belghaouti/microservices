package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/Omar-Belghaouti/microservices/data"
	"github.com/gorilla/mux"
)

type Products struct {
	l *log.Logger
	v *data.Validation
}

func NewProducts(l *log.Logger, v *data.Validation) *Products {
	return &Products{l, v}
}

// ErrInvalidProductPath is an error message when the product path is not valid
var ErrInvalidProductPath = fmt.Errorf("invalid path, path should be /products/[id]")

// GenericError is a generic error message returned by a server
type GenericError struct {
	Message string `json:"message"`
}

// ValidationError is a collection of validation error messages
type ValidationError struct {
	Messages []string `json:"messages"`
}

// swagger:route GET /products/ products getProducts
// Returns a list of products
// responses:
//  200: productsResponse

// GetProducts returns the products from the data store
func (p *Products) GetProducts(w http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle GET Products")
	w.Header().Add("Content-Type", "application/json")
	lp := data.GetProducts()
	err := data.ToJSON(lp, w)
	if err != nil {
		http.Error(w, "Unable to marshal json", http.StatusInternalServerError)
		return
	}
}

// swagger:route GET /products/{id} products getProduct
// Return a product from the database
// responses:
//	200: productResponse
//	404: errorResponse

// GetProduct handles GET requests
func (p *Products) GetProduct(w http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle GET Product")
	w.Header().Add("Content-Type", "application/json")

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Unable to convert id", http.StatusBadRequest)
		return
	}

	prod, err := data.GetProduct(id)
	if err != nil {
		if err == data.ErrProductNotFound {
			w.WriteHeader(http.StatusNotFound)
			data.ToJSON(&GenericError{Message: err.Error()}, w)
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		data.ToJSON(&GenericError{Message: err.Error()}, w)
		return
	}

	err = data.ToJSON(prod, w)
	if err != nil {
		http.Error(w, "Unable to marshal json", http.StatusInternalServerError)
		return
	}
}

// swagger:route POST /products/ products addProduct
// Create a new product
//
// responses:
//	200: productResponse
//  422: errorValidation
//  501: errorResponse

// AddProduct handles POST requests to add new products
func (p *Products) AddProduct(w http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POST Product")
	w.Header().Add("Content-Type", "application/json")

	prod := r.Context().Value(ProductKey{}).(*data.Product)

	data.AddProduct(prod)
}

// swagger:route PUT /products/{id} products updateProduct
// Update a products details
//
// responses:
//	201: noContentResponse
//  404: errorResponse
//  422: errorValidation

// UpdateProduct handles PUT requests to update products
func (p *Products) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle PUT Product")
	w.Header().Add("Content-Type", "application/json")

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Unable to convert id", http.StatusBadRequest)
		return
	}

	prod := r.Context().Value(ProductKey{}).(*data.Product)

	err = data.UpdateProduct(id, prod)
	if err != nil {
		if err == data.ErrProductNotFound {
			w.WriteHeader(http.StatusNotFound)
			data.ToJSON(&GenericError{Message: err.Error()}, w)
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		data.ToJSON(&GenericError{Message: err.Error()}, w)
		return
	}

}

// swagger:route DELETE /products/{id} products deleteProduct
// Update a products details
//
// responses:
//	201: noContentResponse
//  404: errorResponse
//  501: errorResponse

// Delete handles DELETE requests and removes items from the database
func (p *Products) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle DELETE Product")
	w.Header().Add("Content-Type", "application/json")

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Unable to convert id", http.StatusBadRequest)
		return
	}

	err = data.DeleteProduct(id)
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
	err := data.FromJSON(prod, r.Body)
	if err != nil {
		return nil, err
	}
	return prod, nil
}

type ProductKey struct{}

func (p *Products) MiddlewareProductValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")

		prod, err := p.getProductFromBody(w, r)
		if err != nil {
			p.l.Println("[Error] deserializing product", err)
			w.WriteHeader(http.StatusBadRequest)
			data.ToJSON(&GenericError{Message: err.Error()}, w)
			return
		}

		// validate product
		errs := p.v.Validate(prod)
		if errs != nil {
			p.l.Println("[Error] validating product", err)
			w.WriteHeader(http.StatusUnprocessableEntity)
			data.ToJSON(&GenericError{Message: err.Error()}, w)
			return
		}

		ctx := context.WithValue(r.Context(), ProductKey{}, prod)
		req := r.WithContext(ctx)
		next.ServeHTTP(w, req)
	})
}
