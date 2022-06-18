package models

import "time"

type ingredient struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Category    *category `json:"category"`
	Quantity    float64   `json:"quantity"`
	Description string    `json:"desc"`
	CreateOn    string    `json:"-"`
	UpdateOn    string    `json:"-"`
}

func getIngredients() []*ingredient {
	return ingredientList
}

var ingredientList = []*ingredient{
	{
		ID:   1,
		Name: "Chmiel",
		Category: &category{
			ID:   1,
			Name: "Chmiele",
			SubCategory: &category{
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
