package handlers

import (
	"log"
	"net/http"
	"regexp"
	"strconv"

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
		// Use Regular Expression to extract the ID
		rg := regexp.MustCompile(`[0-9]+`)
		d := rg.FindAllStringSubmatch(r.URL.Path, -1)
		// p.l.Println(d)

		// Case : Multiple IDs Match
		if len(d) != 1 {
			p.l.Println("Invalid URI - multiple IDs")
			http.Error(w, "No Match Found!!", http.StatusBadRequest)
			return
		}
		// Case : Multiple Captured in Single Group
		if len(d[0]) < 1 {
			p.l.Println("Invalid URI - more than one capture group")
			http.Error(w, "No Match Found!!", http.StatusBadRequest)
			return
		}

		// Extract the first match of first group
		idString := d[0][0]
		// Convert to Integer
		id, err := strconv.Atoi(idString)
		if err != nil {
			p.l.Println("Invalid URI - ", err)
			http.Error(w, "Invalid URI", http.StatusBadRequest)
			return
		}
		// p.l.Println("ID Received", id)

		// Run update product method on the product
		p.UpdateProduct(id, w, r)
		return
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
