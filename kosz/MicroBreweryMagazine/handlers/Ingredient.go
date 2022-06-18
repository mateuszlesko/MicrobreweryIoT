package handlers

import (
	"log"
	"net/http"
)

type Ingredient struct {
	l *log.Logger
}

func NewIngredient(l *log.Logger) *Ingredient {
	return &Ingredient{l}
}

func (i *Ingredient) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	li := models
	rw.Write([]byte("Available Ingredients"))
}
