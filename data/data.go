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

// create a slice of Product struct Just for ease
type Products []*Product

// Define a method to return in JSON Encoding
func (p *Products) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func GetProducts() Products {
	return productList
}

// Products Slice
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
