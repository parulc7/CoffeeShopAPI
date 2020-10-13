package data

import "testing"

func TestChecksValidation(t *testing.T) {
	p := &Product{
		Name:  "Espresso",
		Price: 10.0,
		SKU:   "abc-adfa-adfdasf",
	}

	err := p.Validate()

	if err != nil {
		t.Fatal(err)
	}
}
