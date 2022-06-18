package handlers

import (
	"log"
	"net/http"

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
		i.getIngredients(rw, r)
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
