package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/mateuszlesko/MicroBreweryIoT/MicroBreweryMagazine2/data"
)

type Category struct {
	l *log.Logger
}

type KeyCategory struct{}

func NewCategory(l *log.Logger) *Category {
	return &Category{l}
}

//get
func (c *Category) GetCategories(rw http.ResponseWriter, r *http.Request) {
	cl, err := data.SelectCategories()
	if err != nil {
		http.Error(rw, "unable to query", http.StatusUnprocessableEntity)
		return
	}
	categoriesBytes, err := json.MarshalIndent(cl, "", "\t")
	if err != nil {
		http.Error(rw, "unable to marshal", http.StatusUnprocessableEntity)
		return
	}
	rw.Header().Set("Content-Type", "application/json")
	rw.Write(categoriesBytes)
}

//get
func (c *Category) GetCategory(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(rw, "unable to encode json", http.StatusBadRequest)
		return
	}
	var category *data.Category
	err, category = data.SelectCategoryWhereID(id)
	if err != nil {
		c.l.Panic(err)
	}
	categoryBytes, err := json.Marshal(category)
	if err != nil {
		http.Error(rw, "unable to marshal", http.StatusUnprocessableEntity)
		return
	}
	rw.Header().Set("Content-Type", "application/json")
	rw.Write(categoryBytes)
}

//delete
func (c *Category) DeleteCategory(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(rw, "unable to decode json", http.StatusBadRequest)
		return
	}
	err, _ = data.SelectCategoryWhereID(id)
	if err != nil {
		http.Error(rw, "no object corresponds in db", http.StatusBadRequest)
		return
	}
	err = data.DeleteCategory(id)
	if err != nil {
		http.Error(rw, "delete was not executed", http.StatusBadRequest)
		return
	}
	rw.Header().Set("Content-Type", "application/json")
	rw.Write([]byte("1"))
}

func (c *Category) UpdateCategory(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(rw, "unable to decode json", http.StatusBadRequest)
		return
	}
	err, _ = data.SelectCategoryWhereID(id)
	if err != nil {
		http.Error(rw, "no object corresponds in db", http.StatusBadGateway)
		return
	}
	category := data.Category{}
	err = json.NewDecoder(r.Body).Decode(&category)
	if err != nil {
		http.Error(rw, "not enable to decode json", http.StatusBadGateway)
	}
	err, _ = data.UpdateCategory(category)
	if err != nil {
		http.Error(rw, "unable to update", http.StatusUnprocessableEntity)
		return
	}
	categoryBytes, err := json.Marshal(category)
	if err != nil {
		http.Error(rw, "unable to marshal", http.StatusUnprocessableEntity)
		return
	}
	rw.Header().Set("Content-Type", "application/json")
	rw.Write(categoryBytes)
}

func (c Category) PostCategory(rw http.ResponseWriter, r *http.Request) {
	var category data.Category
	err := json.NewDecoder(r.Body).Decode(&category)
	if err != nil {
		http.Error(rw, "can not decode value from body", http.StatusBadRequest)
		return
	}
	err, res := data.InsertCategory(category.Category_name)
	if err != nil {
		http.Error(rw, "can not add row", http.StatusBadRequest)
		return
	}
	if err != nil {
		http.Error(rw, "unable to marshal", http.StatusBadRequest)
		return
	}
	rw.Header().Set("Content-Type", "application/json")
	rw.Write([]byte(fmt.Sprintf("%d", res)))
}
