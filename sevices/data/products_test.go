package data

import "testing"

func TestChecksValidation(t *testing.T) {
	p := &Product{
		Name:  "Rick",
		Price: 1.00,
		Sku:   "dfd-343-fdf",
	}

	err := p.Validate()

	if err != nil {
		t.Fatal(err)
	}
}
