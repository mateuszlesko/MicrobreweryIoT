package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/mateuszlesko/MicroBreweryIoT/MicroBreweryMagazine2/data"
)

type KeyIngredient struct{}

type Ingredient struct {
	l *log.Logger
}

func NewIngredient(l *log.Logger) *Ingredient {
	return &Ingredient{l}
}

//get
func (i *Ingredient) GetIngredients(rw http.ResponseWriter, r *http.Request) {
	il, err := data.SelectIngredients()
	if err != nil {
		http.Error(rw, "unable to get data", http.StatusBadRequest)
	}
	ingredientsBytes, err := json.MarshalIndent(il, "", "\t")
	if err != nil {
		http.Error(rw, "unable to encode json", http.StatusBadRequest)
	}
	rw.Header().Set("Content-Type", "application/json")
	rw.Write(ingredientsBytes)
}

func (i *Ingredient) GetIngredient(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(rw, "unable to encode json", http.StatusBadRequest)
	}
	ingredient, err := data.SelectIngredientById(id)
	if err != nil {
		http.Error(rw, "0 record", http.StatusBadRequest)
	}
	ingredientsBytes, err := json.MarshalIndent(ingredient, "", "\t")
	if err != nil {
		http.Error(rw, "unable to encode json", http.StatusBadRequest)
	}
	rw.Header().Set("Content-Type", "application/json")
	rw.Write(ingredientsBytes)
}

//post
func (i *Ingredient) AddIngredient(rw http.ResponseWriter, r *http.Request) {
	ingredient := r.Context().Value(KeyIngredient{}).(data.IngredientVM)
	_, err := data.InsertIngredient(&ingredient)
	if err != nil {
		http.Error(rw, "Unable to add", http.StatusBadRequest)
	}
	rw.Header().Set("Content-Type", "application/json")
}

//put
func (i *Ingredient) UpdateIngredient(rw http.ResponseWriter, r *http.Request) {

	d := r.Context().Value(KeyIngredient{}).(data.IngredientVM)
	err := data.UpdateIngredient(&d)
	if err != nil {
		http.Error(rw, "unable to update", http.StatusBadRequest)
	}

	ingredientDb, err := data.SelectIngredientById(d.Ingredient_id)

	if err != nil {
		http.Error(rw, "unable to get updated data", http.StatusBadRequest)

	}

	ingredientsBytes, err := json.MarshalIndent(ingredientDb, "", "\t")

	if err != nil {
		http.Error(rw, "unable to encode data", http.StatusBadRequest)
	}
	rw.Write(ingredientsBytes)
}

//delete
func (i *Ingredient) DeleteIngredient(rw http.ResponseWriter, r *http.Request) {
	d := mux.Vars(r)
	id, err := strconv.Atoi(d["id"])
	if err != nil && id != 0 {
		http.Error(rw, "unable to get id", http.StatusBadRequest)
	}
	err = data.DeleteIngredient(id)
	if err != nil {
		http.Error(rw, "unable to delete", http.StatusBadRequest)
	}
	rw.Header().Set("Content-Type", "application/json")
}

func (i *Ingredient) MiddlewareIngredientValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {

		ingredient := data.IngredientVM{}
		err := ingredient.FromJSON(r.Body)

		if err != nil {
			http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)
			return
		}

		//validate ingredient data
		err_v := ingredient.Validate()

		if err_v != nil {
			http.Error(rw, "Error validating ingredient", http.StatusBadRequest)
			return
		}
		fmt.Println("MIDDLEWARE")
		//add ingredient to the context
		ctx := context.WithValue(r.Context(), KeyIngredient{}, ingredient)
		req := r.WithContext(ctx)

		//calling the next handler or middleware in the chain
		next.ServeHTTP(rw, req)
	})
}
