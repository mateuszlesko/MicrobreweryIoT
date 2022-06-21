package data

import (
	"encoding/json"
	"fmt"
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

var NotFoundError = fmt.Errorf("NOT FOUND RESOURCE")

func (i *Ingredient) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(i)
}
func (i *Ingredients) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(i) //returns error
}

func (i *Ingredient) FromJSON(r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(i) // pass reference to myself, map json to struct
}

func GetIngredients() Ingredients {
	return ingredientList
}

func AddIngredient(i *Ingredient) {
	i.ID = getNextID()
	ingredientList = append(ingredientList, i)
}

func UpdateIngredient(id int, i *Ingredient) error {
	_, index, err := findIngredient(id)
	if err != nil && index < 0 {
		return err
	}

	i.ID = id
	ingredientList[index] = i
	return nil
}

func RemoveIngredient(id int) error {
	_, index, err := findIngredient(id)
	if err != nil && index < 0 {
		return err
	}
	ingredientList[index] = nil
	return nil
}

func FindIngredient(id int) (*Ingredient, int, error) {
	return findIngredient(id)
}

func getNextID() int {
	return ingredientList[len(ingredientList)-1].ID + 1
}

func findIngredient(id int) (*Ingredient, int, error) {
	for i, ingredient := range ingredientList {
		if ingredient.ID == id {
			return ingredient, i, NotFoundError
		}
	}
	return nil, -1, fmt.Errorf("RESOURCE NOT FOUND")
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
