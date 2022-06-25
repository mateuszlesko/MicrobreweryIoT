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
	"github.com/mateuszlesko/MicroBreweryIoT/MicroBreweryMagazine2/handlers"
)

func main() {

	ch := gorilla_hander.CORS(gorilla_hander.AllowedOrigins([]string{"http://localhost:3000"}))
	l := log.New(os.Stdout, "MicroBreweryMagazineService", log.LstdFlags)
	ih := handlers.NewIngredient(l)

	smux := mux.NewRouter()
	getRouter := smux.Methods(http.MethodGet).Subrouter()
	putRouter := smux.Methods(http.MethodPut).Subrouter()
	postRouter := smux.Methods(http.MethodPost).Subrouter()
	deleteRouter := smux.Methods(http.MethodDelete).Subrouter()

	getRouter.HandleFunc("/ingredients/", ih.GetIngredients)
	getRouter.HandleFunc("/ingredients/{id:[0-9]+}", ih.GetIngredient)

	postRouter.HandleFunc("/ingredients/", ih.AddIngredient)
	postRouter.Use(ih.MiddlewareIngredientValidation)

	putRouter.HandleFunc("/ingredients/{id:[0-9]+}", ih.UpdateIngredient)
	putRouter.Use(ih.MiddlewareIngredientValidation)

	deleteRouter.HandleFunc("/ingredients/{id:[0-9]+}", ih.DeleteIngredient)
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
