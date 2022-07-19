package data

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-playground/validator/v10"
	"github.com/mateuszlesko/MicroBreweryIoT/MicroBreweryMagazine2/helpers"
)

type Ingredient struct {
	Ingredient_Id   int       `json:"id"`
	Ingredient_Name string    `json:"name" validate:"required"`
	Unit            string    `json:"unit" validate:"required,unit"`
	Quantity        float32   `json:"quantity" validate:"required"`
	Description     string    `json:"desc"`
	Category        *Category `json:"category"`
}

type IngredientVM struct {
	Ingredient_id          int     `json:"id"`
	Ingredient_name        string  `json:"name" validate:"required"`
	Ingredient_unit        string  `json:"unit" validate:"required,unit"`
	Ingredient_quantity    float32 `json:"quantity" validate:"required"`
	Ingredient_description string  `json:"desc"`
	Category_id            int     `json:"category"`
}

var NotFoundError = fmt.Errorf("NOT FOUND RESOURCE")

func (i *Ingredient) Validate() error {
	v := validator.New()
	v.RegisterValidation("unit", validateUnit)
	return v.Struct(i)
}

func (i *IngredientVM) Validate() error {
	v := validator.New()
	v.RegisterValidation("unit", validateUnit)
	return v.Struct(i)
}

func validateUnit(fl validator.FieldLevel) bool {
	switch fl.Field().String() {
	case
		"mg",
		"g",
		"dag",
		"kg",
		"t":
		return true
	}
	return false
}

func (i *Ingredient) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(i)
}

func (ivm *IngredientVM) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(ivm)
}

func (i *Ingredient) FromJSON(r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(i) // pass reference to myself, map json to struct
}

func (i *IngredientVM) FromJSON(r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(i) // pass reference to myself, map json to struct
}

func SelectIngredients() ([]Ingredient, error) {
	err, db := helpers.OpenConnection()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	//rows, err := db.Query("SELECT ingredients.ingredient_id,ingredients.ingredient_name,ingredients.unit,ingredients.quantity,ingredients.description FROM ingredients inner join categories on ingredients.category_id = categories.category_id;")
	rows, err := db.Query("SELECT ingredients.ingredient_id,ingredients.ingredient_name,ingredients.unit,ingredients.quantity,categories.category_id,categories.category_name FROM ingredients inner join categories on ingredients.category_id = categories.category_id;")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	il := []Ingredient{}
	var (
		ingredient_id          int
		ingredient_name        string
		ingredient_unit        string
		ingredient_quantity    float32
		ingredient_description string
		category_id            int
		category_name          string
	)
	for rows.Next() {
		fmt.Println(rows.ColumnTypes())
		err = rows.Scan(&ingredient_id, &ingredient_name, &ingredient_unit, &ingredient_quantity, &category_id, &category_name)
		if err != nil {
			rows.Close()
			db.Close()
			return nil, err
		}
		category := &Category{category_id, category_name}
		ingredient := &Ingredient{ingredient_id, ingredient_name, ingredient_unit, ingredient_quantity, ingredient_description, category}
		il = append(il, *ingredient)
	}

	return il, err
}

func SelectIngredientById(id int) (*Ingredient, error) {
	err, db := helpers.OpenConnection()
	if err != nil {
		return nil, err
	}

	var (
		ingredient_id       int
		ingredient_name     string
		ingredient_unit     string
		ingredient_quantity float32
		//ingredient_description string
		category_id   int
		category_name string
	)
	fmt.Println(id)
	if err := db.QueryRow("SELECT ingredients.ingredient_id,ingredients.ingredient_name,ingredients.unit,ingredients.quantity,categories.category_id,categories.category_name FROM ingredients inner join categories on ingredients.category_id = categories.category_id where ingredients.ingredient_id=$1", id).Scan(&ingredient_id, &ingredient_name, &ingredient_unit, &ingredient_quantity, &category_id, &category_name); err != nil {
		return nil, err
	}
	//category := &Category{category_id, category_name}
	ingredient := &Ingredient{ingredient_id, ingredient_name, ingredient_unit, ingredient_quantity, "", &Category{category_id, category_name}}
	defer db.Close()
	return ingredient, err
}

func InsertIngredient(i *IngredientVM) (int, error) {
	err, db := helpers.OpenConnection()
	if err != nil {
		return -1, err
	}
	//"SELECT ingredient_id,ingredient_name,unit,quantity,category_id
	smt, err := db.Prepare(`insert into ingredients(ingredient_id,ingredient_name,unit,quantity,category_id) values($1,$2,$3,$4)`)
	if err != nil {
		return -1, err
	}
	_, err = smt.Exec(i.Ingredient_name, i.Ingredient_unit, i.Ingredient_quantity, i.Ingredient_description, i.Category_id)
	if err != nil {
		return -1, err
	}
	if err != nil {
		return -1, err
	}
	defer smt.Close()
	defer db.Close()
	return 1, nil
}
