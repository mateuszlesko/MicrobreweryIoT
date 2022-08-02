package data

import (
	"encoding/json"
	"io"
	"time"

	"github.com/mateuszlesko/MicroBreweryIoT/MicroBreweryMagazine2/helpers"
)

type IngredientCategory struct {
	Category_id         int       `json:"id"`
	Category_name       string    `json:"name"`
	Category_created_at time.Time `json:"createdAt"`
}

func CreateIngredientCategory(id int, name string) *IngredientCategory {
	c := IngredientCategory{}
	c.Category_id = id
	c.Category_name = name
	return &c
}

/// ***********CATEGORY********************
func (c *IngredientCategory) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(c)
}
func ToJSON(c []IngredientCategory, w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(c)
}
func (c *IngredientCategory) FromJSON(r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(c) // pass reference to myself, map json to struct
}

/// **********CATEGORY DB LOGIC********************

func SelectCategories() ([]IngredientCategory, error) {
	err, db := helpers.OpenConnection()

	if err != nil {
		return nil, err
	}

	rows, err := db.Query("SELECT * FROM ingredient_category")
	if err != nil {
		return nil, err
	}

	var categories []IngredientCategory

	for rows.Next() {
		var category IngredientCategory
		if err := rows.Scan(&category.Category_id, &category.Category_name, &category.Category_created_at); err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}
	defer rows.Close()
	defer db.Close()
	return categories, nil
}

func SelectCategoryWhereID(id int) (*IngredientCategory, error) {
	err, db := helpers.OpenConnection()
	defer db.Close()
	if err != nil {
		return nil, err
	}
	var category IngredientCategory

	if err := db.QueryRow("SELECT * FROM ingredient_category WHERE id=$1;", id).Scan(&category.Category_id, &category.Category_name, &category.Category_created_at); err != nil {
		return nil, err
	}

	return &category, nil
}

func UpdateCategory(category IngredientCategory) (*IngredientCategory, error) {
	err, db := helpers.OpenConnection()
	if err != nil {
		return nil, err
	}
	smt, err := db.Prepare(`update ingredient_category set category_name=$1 where id=$2`)
	if err != nil {
		return nil, err
	}
	if _, err := smt.Exec(category.Category_name, category.Category_id); err != nil {
		return nil, err
	}
	defer smt.Close()
	defer db.Close()
	return &category, nil
}

func InsertCategory(name string) error {
	err, db := helpers.OpenConnection()
	if err != nil {
		return err
	}
	smt, err := db.Prepare(`insert into ingredient_category(category_name,created_at) values($1,NOW());`)
	if err != nil {
		return err
	}
	_, err = smt.Exec(name)
	if err != nil {
		return err
	}
	if err != nil {
		return err
	}
	defer smt.Close()
	defer db.Close()
	return nil
}

func DeleteCategory(id int) error {
	err, db := helpers.OpenConnection()
	defer db.Close()
	if err != nil {
		return err
	}

	//make sure of data integration - remove all connection to data that we want to delete
	smt, err := db.Prepare(`update ingredient set category_name=null where ingredient_category_id=$1`)
	if err != nil {
		return err
	}
	if _, err = smt.Exec(id); err != nil {
		return err
	}

	smt, err = db.Prepare(`delete from ingredient_category where id=$1;`)
	if err != nil {
		return err
	}
	defer smt.Close()
	if err != nil {
		return err
	}
	if _, err := smt.Exec(id); err != nil {
		return err
	}

	return nil
}
