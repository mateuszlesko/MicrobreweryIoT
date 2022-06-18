package data

import (
	"encoding/json"
	"io"
	"time"
)

type Ingredient struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Category    *Category `json:"category"`
	Quantity    float64   `json:"quantity"`
	Description string    `json:"desc"`
	CreateOn    string    `json:"-"`
	UpdateOn    string    `json:"-"`
}

type Ingredients []*Ingredient

func (i *Ingredients) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(i) //returns error
}

func GetIngredients() Ingredients {
	return ingredientList
}

var ingredientList = []*Ingredient{
	{
		ID:   1,
		Name: "Chmiel",
		Category: &Category{
			ID:   1,
			Name: "Chmiele",
			SubCategory: &Category{
				ID:   2,
				Name: "Chmiele jasne",
			},
		},
		Quantity:    100000.0,
		Description: "Jasny chmiel",
		CreateOn:    time.Now().UTC().String(),
		UpdateOn:    time.Now().UTC().String(),
	},
}
