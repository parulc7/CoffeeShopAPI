package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/parulc7/CoffeeShopAPI/handlers"
)

func main() {
	// Global Uniform Logger
	l := log.New(os.Stdout, "coffee-shop", log.LstdFlags)
	// Passing the global logger to the handler
	productsHandler := handlers.NewProducts(l)
	// Create a custom servemux
	sm := http.NewServeMux()

	// Server Config here
	s := &http.Server{
		Addr:         ":8080",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	// Listen and serve concurrently while waiting for an interrupt
	go func() {
		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()

	// Map the routes in servemux and start server
	sm.Handle("/", productsHandler)

	// Graceful Shutdown
	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Kill)
	signal.Notify(sigChan, os.Interrupt)

	sig := <-sigChan
	l.Println("Received a termination signal :: ", sig)
	ctx, _ := context.WithTimeout(context.TODO(), 30*time.Second)
	s.Shutdown(ctx)
}
