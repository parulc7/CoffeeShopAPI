package handlers

import (
	"log"
	"net/http"
	"regexp"

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
	if r.Method == http.MethodPost {
		p.AddProduct(w, r)
		return
	}
	if r.Method == http.MethodPut {
		// TODO: GET ID AND HANDLE THE PUT REQUEST

		// Use Regular Expression to extract the ID
		r := regexp.MustCompile(`/(0-9)+`)
		d := r.FindAllStringSubmatch(r.URL.Path)
	}
	// Catch all
	w.WriteHeader(http.StatusMethodNotAllowed)
}

// GET Request Handler Function
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

// POST Request Handler Function
func (p *Product) AddProduct(w http.ResponseWriter, r *http.Request) {
	p.l.Println("POST Request Received for Products Model!!")
	prod := &data.Product{}
	err := prod.ToModel(r.Body)
	if err != nil {
		http.Error(w, "Error while posting data!!", http.StatusBadRequest)
		// p.l.Println(err)
	}
	// p.l.Println(prod)
	data.AddProduct(prod)
}

// PUT Request Handler Function
func (p *Product) UpdateProduct(id int, w http.ResponseWriter, r *http.Request) {
	p.l.Println("PUT Request Received for Products Model!!")
	prod := &data.Product{}
	err := prod.ToModel(r.Body)
	if err != nil {
		http.Error(w, "Error while posting data!!", http.StatusBadRequest)
		// p.l.Println(err)
	}
	// p.l.Println(prod)

	// Call Update method
	err = data.UpdateProduct(id, prod)

	// If Product not found in DB
	if err == data.ErrorProductNotFound {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}
	// If any other error
	if err != nil {
		http.Error(w, "Error while Updating Product", http.StatusInternalServerError)
		return
	}
}
