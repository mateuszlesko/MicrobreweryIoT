package handlers

import (
	"log"
	"net/http"
	"regexp"
	"strconv"

	"github.com/mateuszlesko/MicroBreweryIoT/MicroBreweryMagazine2/data"
)

type Ingredient struct {
	l *log.Logger
}

func NewIngredient(l *log.Logger) *Ingredient {
	return &Ingredient{l}
}

func (i *Ingredient) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		rr := regexp.MustCompile(`/([0-9]+)`)
		g := rr.FindAllStringSubmatch(r.URL.Path, -1)
		if len(g) > 0 {
			if len(g[0]) != 2 {
				http.Error(rw, "Invalid URL", http.StatusBadRequest)
				return
			}
			idString := g[0][1]
			id, _ := strconv.Atoi(idString)
			ingredient, _, _ := data.FindIngredient(id)
			ingredient.ToJSON(rw)
			return
		}
		i.getIngredients(rw, r)
		return
	}
	if r.Method == http.MethodPost {
		i.addIngredient(rw, r)
		return
	}
	if r.Method == http.MethodPut {
		rr := regexp.MustCompile(`/([0-9]+)`)
		g := rr.FindAllStringSubmatch(r.URL.Path, -1)
		if len(g) != 1 {
			http.Error(rw, "Invalid URL", http.StatusBadRequest)
			return
		}
		if len(g[0]) != 2 {
			http.Error(rw, "Invalid URL", http.StatusBadRequest)
			return
		}
		idString := g[0][1]
		id, _ := strconv.Atoi(idString)
		i.l.Printf("%d", id)
		i.updateIngredient(id, rw, r)
		return
	}
	if r.Method == http.MethodDelete {
		rr := regexp.MustCompile(`/([0-9]+)`)
		g := rr.FindAllStringSubmatch(r.URL.Path, -1)
		if len(g) != 1 {
			http.Error(rw, "Invalid URL", http.StatusBadRequest)
			return
		}
		if len(g[0]) != 2 {
			http.Error(rw, "Invalid URL", http.StatusBadRequest)
			return
		}
		idString := g[0][1]
		id, _ := strconv.Atoi(idString)
		i.l.Printf("%d", id)
		i.deleteIngredient(id, rw, r)
		return
	}
	rw.WriteHeader(http.StatusMethodNotAllowed)
}

//get
func (i *Ingredient) getIngredients(rw http.ResponseWriter, r *http.Request) {
	li := data.GetIngredients()
	err := li.ToJSON(rw)
	if err != nil {
		http.Error(rw, "unable to encode json", http.StatusInternalServerError)
	}
}

//post
func (i *Ingredient) addIngredient(rw http.ResponseWriter, r *http.Request) {
	ingredient := &data.Ingredient{}
	err := ingredient.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "unable to unmarshal json", http.StatusBadRequest)
	}
	data.AddIngredient(ingredient)
}

//put
func (i *Ingredient) updateIngredient(id int, rw http.ResponseWriter, r *http.Request) {
	ingredient := &data.Ingredient{}
	err := ingredient.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "unable to unmarshal json", http.StatusBadRequest)
	}
	err = data.UpdateIngredient(id, ingredient)
	if err != nil {
		data.AddIngredient(ingredient)
	}
}

//delete
func (i *Ingredient) deleteIngredient(id int, rw http.ResponseWriter, r *http.Request) {
	err := data.RemoveIngredient(id)
	if err != nil {
		http.Error(rw, "unabable to do this operation", http.StatusBadRequest)
	}
}
