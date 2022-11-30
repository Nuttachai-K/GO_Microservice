package data

import "testing"

func TestCheckValidation(t *testing.T) {
	p := &Product{
		Name:  "Korn",
		Price: 1.00,
		SKU:   "abs-abc-def",
	}

	err := Validate(p)

	if err != nil {
		t.Fatal(err)
	}
}
