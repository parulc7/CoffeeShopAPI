package data

import (
	"encoding/json"
	"io"
	"time"
)

// Product Model
type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	SKU         string  `json:"sku"`
	CreatedOn   string  `json:"-"`
	UpdatedOn   string  `json:"lastUpdated"`
	DeletedOn   string  `json:"-"`
}

var ErrorProductNotFound = fmt.Error("Product Not Found!!")

// create a slice of Product struct Just for ease
type Products []*Product

// Define a method to return in JSON Encoding - GET
func (p *Products) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

// Define a method to convert JSON data to our model - POST/PUT
func (p *Product) ToModel(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(p)
}

// Get the list of products
func GetProducts() Products {
	return productList
}

// Add to the list of Products
func AddProduct(p *Product) {
	// Generat ID for the product
	p.ID = generateID()
	// Add to the data
	productList = append(productList, p)
}

// Helper function to generate ID
func generateID() int {
	id := productList[len(productList)-1].ID
	return id + 1
}

func UpdateProduct(id int, p *Product) error {
	// Find Product
	i, err := findProduct(id)
	if err != nil {
		return err
	}
	// Update the product
	p.ID = id
	productList[i] = p
	return nil
}

// Helper Function to find the product to be updated
func findProduct(id int) (int, error) {
	for i, p := range productList {
		if p.ID == id {
			return i, nil
		}
	}
	return -1, ErrorProductNotFound
}

// Products Slice i.e. the Data
var productList = []*Product{
	&Product{
		ID:          1,
		Name:        "Cafe Latte",
		Description: "A rich brewed black coffee",
		Price:       120.00,
		SKU:         "cl1",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
		DeletedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          2,
		Name:        "Espresso",
		Description: "A rich brewed milk coffee",
		Price:       190.00,
		SKU:         "es1",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
		DeletedOn:   time.Now().UTC().String(),
	},
}
