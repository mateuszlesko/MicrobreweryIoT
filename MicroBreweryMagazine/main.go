package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	gorilla_hander "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/mateuszlesko/MicroBreweryIoT/MicroBreweryMagazine2/handlers"
	"github.com/mateuszlesko/MicroBreweryIoT/MicroBreweryMagazine2/helpers"
)

func main() {
	ch := gorilla_hander.CORS(gorilla_hander.AllowedOrigins([]string{"http://localhost:3000"}))
	l := log.New(os.Stdout, "MicroBreweryMagazineService", log.LstdFlags)
	var err error
	err, pgsql := helpers.OpenConnection()
	if err != nil {
		log.Panic(err)
	}
	err = pgsql.Ping()
	if err != nil {
		panic(err)
	}

	cgh := handlers.NewCategory(l, &pgsql)
	ih := handlers.NewIngredient(l)

	smux := mux.NewRouter()
	getIngredient := smux.Methods(http.MethodGet).Subrouter()
	putIngredient := smux.Methods(http.MethodPut).Subrouter()
	postIngredient := smux.Methods(http.MethodPost).Subrouter()
	deleteIngredient := smux.Methods(http.MethodDelete).Subrouter()

	getIngredient.HandleFunc("/ingredients/", ih.GetIngredients)
	getIngredient.HandleFunc("/ingredients/{id:[0-9]+}", ih.GetIngredient)

	postIngredient.HandleFunc("/ingredients/", ih.AddIngredient)
	postIngredient.Use(ih.MiddlewareIngredientValidation)

	putIngredient.HandleFunc("/ingredients/{id:[0-9]+}", ih.UpdateIngredient)
	putIngredient.Use(ih.MiddlewareIngredientValidation)

	deleteIngredient.HandleFunc("/ingredients/{id:[0-9]+}", ih.DeleteIngredient)

	getCategory := smux.Methods(http.MethodGet).Subrouter()
	putCategory := smux.Methods(http.MethodPut).Subrouter()
	postCategory := smux.Methods(http.MethodPost).Subrouter()
	deleteCategory := smux.Methods(http.MethodDelete).Subrouter()

	getCategory.HandleFunc("/categories/", cgh.GetCategories)
	getCategory.HandleFunc("/category/{id:[0-9]+}", cgh.GetCategory)
	postCategory.HandleFunc("/category/", cgh.InsertCategory)
	putCategory.HandleFunc("/category/{id:[0-9]+}", cgh.UpdateCategory)
	deleteCategory.HandleFunc("/category/{id:[0-9]+}", cgh.DeleteCategory)
	fmt.Println("Server is listening on :6660")

	s := &http.Server{
		Addr:              ":6660",
		Handler:           ch(smux),
		TLSConfig:         &tls.Config{},
		ReadTimeout:       2 * time.Second,
		ReadHeaderTimeout: 2 * time.Second,
		WriteTimeout:      2 * time.Second,
		IdleTimeout:       120 * time.Second,
		MaxHeaderBytes:    16,
		TLSNextProto:      map[string]func(*http.Server, *tls.Conn, http.Handler){},
		ConnState: func(net.Conn, http.ConnState) {
		},
		ErrorLog: l,
	}
	defer pgsql.Close()
	//async doing serve http server
	go func() {
		err := s.ListenAndServe()
		if err == nil {
			l.Fatal(err)
		}
	}()

	//serwer waits unitl everyone finish and does not take any new request, after that it will peacfully shut down.
	sChan := make(chan os.Signal, 1)
	signal.Notify(sChan, os.Interrupt)
	signal.Notify(sChan, syscall.SIGTERM)
	sig := <-sChan
	l.Println("Recieved terminate, graceful shutdown", sig)
	tc, _ := context.WithTimeout(context.Background(), 2*time.Second)
	s.Shutdown(tc)
}

func OpenConnection() {
	panic("unimplemented")
}
