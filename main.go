package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/franciscofferraz/product-api/handlers"
	"github.com/gorilla/mux"
)

func main() {

	l := log.New(os.Stdout, "product-api", log.LstdFlags)

	productsHandler := handlers.NewProducts(l)

	serveMux := mux.NewRouter()

	// GET Routes
	getRouter := serveMux.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/products", productsHandler.GetProducts)

	// PUT Routes
	putRouter := serveMux.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc("/products/{id:[0-9]+}", productsHandler.UpdateProduct)
	putRouter.Use(productsHandler.MiddlewareProductValidation)

	// POST Routes
	postRouter := serveMux.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/products/new", productsHandler.AddProduct)
	postRouter.Use(productsHandler.MiddlewareProductValidation)

	server := &http.Server{
		Addr:         ":9090",
		Handler:      serveMux,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()

	signalChannel := make(chan os.Signal)
	signal.Notify(signalChannel, os.Interrupt)
	signal.Notify(signalChannel, os.Kill)

	sig := <-signalChannel
	l.Println("Received terminate, shutdown", sig)

	timeoutContext, _ := context.WithTimeout(context.Background(), 30*time.Second)
	server.Shutdown(timeoutContext)
}
