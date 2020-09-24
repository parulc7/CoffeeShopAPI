package handlers

import (
	"log"
	"net/http"

	"github.com/parulc7/CoffeeShopAPI/data"
)

// Handler Type to take global logger instance
type Product struct {
	l *log.Logger
}

// Convert to Handler type
func NewProducts(l *log.Logger) *Product {
	return &Product{l}
}

// Server HTTP Method of the Handler Interface
func (p *Product) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// If we get a GET Request, return the products
	if r.Method == http.MethodGet {
		p.GetProducts(w, r)
		return
	}
	// Catch all
	w.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *Product) GetProducts(w http.ResponseWriter, r *http.Request) {
	// prod := []Product{}
	// err := json.NewDecoder(r.Body).Decode(&prod)
	// if err != nil {
	// 	http.Error(w, "Error while reading data from request!!\n", http.StatusBadRequest)
	// 	return
	// }

	// Set response content type to json
	w.Header().Set("Content-Type", "application/json")

	// Get Products data from DB
	productsList := data.GetProducts()
	// data, err := json.Marshal(productsList)
	// if err != nil {
	// 	http.Error(w, "Error in reading data from DB", http.StatusInternalServerError)
	// 	return
	// }
	// p.l.Println(data)
	// w.Write(data)

	// Return as JSON by using Encoder/Marshall
	err := productsList.ToJSON(w)
	// Error Handling
	if err != nil {
		http.Error(w, "Error while getting data!!", http.StatusInternalServerError)
		return
	}
}
