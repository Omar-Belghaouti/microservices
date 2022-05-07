package data

import "testing"

func TestCheckValidation(t *testing.T) {
	p := &Product{
		Name:  "Test",
		Price: 1,
		SKU:   "abc-deds-dsds",
	}

	err := p.Validate()
	if err != nil {
		t.Fatal(err)
	}
}
