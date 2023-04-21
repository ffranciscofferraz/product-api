package data

import "testing"

func TestChecksValidation(t *testing.T) {
	p := &Product{
		Name:  "Latte",
		Price: 2.45,
		SKU:   "abc-abf-fdg",
	}

	err := p.Validate()

	if err != nil {
		t.Fatal(err)
	}
}
