package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/mateuszlesko/MicroBreweryIoT/MicroBreweryMagazine2/data"
)

type Category struct {
	l   *log.Logger
	sql *sql.DB
}

type KeyCategory struct{}

func NewCategory(l *log.Logger, s *sql.DB) *Category {
	return &Category{l, s}
}

//get
func (c *Category) GetCategories(rw http.ResponseWriter, r *http.Request) {
	cl, err := data.SelectCategories()
	if err != nil {
		http.Error(rw, "unable to query", http.StatusUnprocessableEntity)
	}
	categoriesBytes, err := json.MarshalIndent(cl, "", "\t")
	if err != nil {
		http.Error(rw, "unable to marshal", http.StatusUnprocessableEntity)
	}
	rw.Header().Set("Content-Type", "application/json")
	rw.Write(categoriesBytes)
}

func (c *Category) GetCategory(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(rw, "unable to encode json", http.StatusBadRequest)
	}
	var category *data.Category
	err, category = data.SelectCategoryWhereID(id)
	if err != nil {
		c.l.Panic(err)
	}
	categoryBytes, err := json.Marshal(category)
	if err != nil {
		http.Error(rw, "unable to marshal", http.StatusUnprocessableEntity)
	}
	rw.Header().Set("Content-Type", "application/json")
	rw.Write(categoryBytes)
}

func SelectCategory(id int, dB *sql.DB) {
	panic("unimplemented")
}

//post
func (c *Category) InsertCategory(rw http.ResponseWriter, r *http.Request) {

}

//delete
func (c *Category) DeleteCategory(rw http.ResponseWriter, r *http.Request) {

}

func (c *Category) UpdateCategory(rw http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(rw, "unable to encode json", http.StatusBadRequest)
	}
	err, _ = data.SelectCategoryWhereID(id)
	if err != nil {
		http.Error(rw, "no object corresponds in db", http.StatusBadGateway)
	}
	category := data.Category{}
	err = json.NewDecoder(r.Body).Decode(&category)
	if err != nil {
		http.Error(rw, "not enable to decode json", http.StatusBadGateway)
	}
	// err, category = data.UpdateCategory()
	// if err != nil {
	// 	http.Error(rw, "unable to update", http.StatusUnprocessableEntity)
	// }
	categoryBytes, err := json.Marshal(category)
	if err != nil {
		http.Error(rw, "unable to marshal", http.StatusUnprocessableEntity)
	}
	rw.Header().Set("Content-Type", "application/json")
	rw.Write(categoryBytes)

}
