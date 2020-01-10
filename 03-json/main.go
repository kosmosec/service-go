package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

//Create a type Product with fields for Name, Cost, and Quantity
//Rename the Echo handler to ListProducts
//Create a slice of Product values with some dummy data.
//Marshal the slice to JSON and write it to the client.
//Use w.WriteHeader to explicitly set the response status code.
//Include the Content-Type header so clients understand the response. w.Header().Set("Content-Type", "application/json; charset=utf-8")
//See what happens when a nil slice is provided.

func main() {

	api := &http.Server{
		Addr:         "localhost:8080",
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		Handler:      http.HandlerFunc(listProducts),
	}

	serverError := make(chan error, 1)
	log.Printf("Start listening on addr %s\n", api.Addr)
	go func() {
		serverError <- api.ListenAndServe()
	}()

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

	select {
	case err := <-serverError:
		log.Fatalf("Fatal from server %s", err)
	case <-shutdown:
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		// zamyka wszystkie polaczenia
		err := api.Shutdown(ctx)
		if err != nil {
			// imidately close
			err = api.Close()
		}
		if err != nil {
			log.Fatalf("main : could not stop server gracefully : %v", err)
		}
	}

}

func listProducts(w http.ResponseWriter, r *http.Request) {
	products := []product{
		{Name: "iPhone", Cost: "w ciuld", Quantity: "dobra"},
		{Name: "mac", Cost: "w ciuld", Quantity: "dobra"},
		{Name: "tablet", Cost: "w ciuld", Quantity: "dobra"},
		{Name: "izegarek", Cost: "w ciuld", Quantity: "dobra"},
	}

	body, err := json.Marshal(products)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w ,"something goes wrong: %s", err)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	if _, err := w.Write(body); err != nil {
		log.Println("error writing result", err)
	}
}

type product struct {
	Name     string `json:"name"`
	Cost     string `json:"cost"`
	Quantity string `json:"quantity"`
}
