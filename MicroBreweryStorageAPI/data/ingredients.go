package data

import (
	"encoding/json"
	"fmt"
	"io"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/mateuszlesko/MicroBreweryIoT/MicroBreweryMagazine2/helpers"
)

type Ingredient struct {
	Ingredient_Id   int                 `json:"id"`
	Ingredient_Name string              `json:"name" validate:"required"`
	Unit            string              `json:"unit" validate:"required,unit"`
	Quantity        float32             `json:"quantity" validate:"required"`
	CreatedAt       time.Time           `json:"created_at"`
	Category        *IngredientCategory `json:"category"`
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
	rows, err := db.Query("select id, ingredient_name, unit, quantity,created_at from ingredient;")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	il := []Ingredient{}
	var (
		ingredient_id         int
		ingredient_name       string
		ingredient_unit       string
		ingredient_quantity   float32
		ingredient_created_at time.Time
		category_id           int
		category_name         string
	)
	for rows.Next() {
		err = rows.Scan(&ingredient_id, &ingredient_name, &ingredient_unit, &ingredient_quantity, &ingredient_created_at)
		if err != nil {
			rows.Close()
			db.Close()
			return nil, err
		}
		category := CreateIngredientCategory(category_id, category_name)
		//category := &IngredientCategory{category_id, category_name}
		ingredient := &Ingredient{ingredient_id, ingredient_name, ingredient_unit, ingredient_quantity, ingredient_created_at, category}
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
		ingredient_id         int
		ingredient_name       string
		ingredient_unit       string
		ingredient_quantity   float32
		ingredient_created_at time.Time
		category_id           int
		category_name         string
	)

	if err := db.QueryRow("SELECT ingredient.id,ingredient.ingredient_name,ingredient.unit,ingredient.quantity,ingredient.created_at,ingredient_category.id,ingredient_category.category_name FROM ingredient inner join ingredient_category on ingredient.ingredient_category_id = ingredient_category.id where ingredient.id=$1;", id).Scan(&ingredient_id, &ingredient_name, &ingredient_unit, &ingredient_quantity, &ingredient_created_at, &category_id, &category_name); err != nil {
		return nil, err
	}

	ingredient := &Ingredient{ingredient_id, ingredient_name, ingredient_unit, ingredient_quantity, ingredient_created_at, CreateIngredientCategory(category_id, category_name)}
	defer db.Close()
	return ingredient, err
}

func InsertIngredient(i *IngredientVM) (int, error) {
	err, db := helpers.OpenConnection()
	if err != nil {
		return -1, err
	}
	defer db.Close()

	smt, err := db.Prepare("insert into ingredient(ingredient_name,unit,quantity,category_id,created_at) values($1,$2,$3,$4,NOW());")
	if err != nil {
		return -1, err
	}
	defer smt.Close()

	_, err = smt.Exec(i.Ingredient_name, i.Ingredient_unit, i.Ingredient_quantity, i.Category_id)
	if err != nil {
		return -1, err
	}
	return 1, nil
}

func UpdateIngredient(i *IngredientVM) error {
	err, db := helpers.OpenConnection()
	if err != nil {
		return err
	}
	defer db.Close()

	smt, err := db.Prepare("update ingredient set ingredient_name=$1,unit=$2,quantity=CAST($3 AS NUMERIC),category_id=$4 where id=$5")

	if err != nil {
		fmt.Println("prepare err")
		return err
	}

	defer smt.Close()

	if _, err := smt.Exec(i.Ingredient_name, i.Ingredient_unit, i.Ingredient_quantity, i.Category_id, i.Ingredient_id); err != nil {
		fmt.Println("exec error", i.Ingredient_quantity)
		return err
	}

	return nil
}

func DeleteIngredient(id int) error {
	_, err := SelectIngredientById(id)
	if err != nil {
		return err
	}
	err, db := helpers.OpenConnection()
	if err != nil {
		return err
	}
	defer db.Close()
	smt, err := db.Prepare(`delete from ingredient where id=$1;`)

	if err != nil {
		return err
	}
	defer smt.Close()

	if _, err := smt.Exec(id); err != nil {
		return err
	}
	return nil
}

func CheckStock(id int, needStock float32, unit string) (int, error) {
	err, db := helpers.OpenConnection()
	if err != nil {
		return 0, err
	}
	defer db.Close()
	var res int

	if err := db.QueryRow("Select CASE WHEN quantity > $1 and unit = $2 THEN 1 ELSE 0 END AS res from ingredient where id = $3;", needStock, unit, id).Scan(&res); err != nil {
		return 0, err
	}
	return res, nil
}
