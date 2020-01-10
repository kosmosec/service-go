package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	log.Println("Main: Started")
	defer log.Println("Main: Completed")

	h := http.TimeoutHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(5 *time.Second)
		fmt.Fprintf(w, "hello you requested %s %s", r.Method, r.URL.Path)
	}), 2 *time.Second, "timeout gw")

	api := &http.Server{
		Addr: "localhost:8080",
		//Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//	time.Sleep(5 *time.Second)
		//	fmt.Fprintf(w, "hello you requested %s %s", r.Method, r.URL.Path)
		//}),
		Handler: h,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 15 * time.Second,
	}

	// Make a channel to listen for errors coming from the listener. Use a
	// buffered channel so the goroutine can exit if we don't collect this error
	serverErrors := make(chan error, 1)

	go func() {
		log.Printf("main: API listening on %s", api.Addr)
		serverErrors <- api.ListenAndServe()
	}()

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

	select {
	case err := <-serverErrors:
		log.Fatalf("error: listening and serving: %s", err)
	case <-shutdown:
		log.Println("Main: start shutdown")
		// Given outstanding requests a deadling for completion.
		const timeout = 5 * time.Second
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()

		// Asking listener to shutdown and load shed.
		err := api.Shutdown(ctx)
		if err != nil {
			log.Printf("main: Graceful shutdown did not complete in %v : %v", timeout, err)
			err = api.Close()
		}

		if err != nil {
			log.Fatalf("main : could not stop server gracefully : %v", err)
		}
	}
}
