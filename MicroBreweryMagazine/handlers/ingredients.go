package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/mateuszlesko/MicroBreweryIoT/MicroBreweryMagazine2/data"
)

type Ingredient struct {
	l *log.Logger
}

func NewIngredient(l *log.Logger) *Ingredient {
	return &Ingredient{l}
}

//get
func (i *Ingredient) GetIngredients(rw http.ResponseWriter, r *http.Request) {
	li := data.GetIngredients()
	err := li.ToJSON(rw)
	if err != nil {
		http.Error(rw, "unable to encode json", http.StatusInternalServerError)
	}
}

func (i *Ingredient) GetIngredient(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(rw, "unable to do operation", http.StatusBadRequest)
	}
	ingredient, _, _ := data.FindIngredient(id)
	err = ingredient.ToJSON(rw)
	if err != nil {
		http.Error(rw, "unable to do operation", http.StatusBadRequest)
	}
}

//post
func (i *Ingredient) AddIngredient(rw http.ResponseWriter, r *http.Request) {
	ingredient := &data.Ingredient{}
	err := ingredient.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "unable to unmarshal json", http.StatusBadRequest)
	}
	data.AddIngredient(ingredient)
}

//put
func (i *Ingredient) UpdateIngredient(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		http.Error(rw, "unable to do this operation", http.StatusBadRequest)
	}

	ingredient := &data.Ingredient{}
	err = ingredient.FromJSON(r.Body)

	if err != nil {
		http.Error(rw, "unable to unmarshal json", http.StatusBadRequest)
	}
	err = data.UpdateIngredient(id, ingredient)
	if err != nil {
		http.Error(rw, "unable to do this operation", http.StatusBadRequest)
	}
}

//delete
func (i *Ingredient) DeleteIngredient(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(rw, "unable to do this operation", http.StatusBadRequest)
	}
	err = data.RemoveIngredient(id)
	if err != nil {
		http.Error(rw, "unabable to do this operation", http.StatusBadRequest)
	}
}
