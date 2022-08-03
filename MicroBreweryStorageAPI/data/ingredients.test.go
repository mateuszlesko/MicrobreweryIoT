package data

import (
	"testing"
)

func TestChecksValidation(t *testing.T) {
	i := &Ingredient{}
	err := i.Validate()

	if err != nil {
		t.Fatal(err)
	}
}
