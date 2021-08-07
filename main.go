package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/borkod/Continuous-Delivery-in-Go-Translation-API/handlers"
	"github.com/borkod/Continuous-Delivery-in-Go-Translation-API/handlers/rest"
)

func main() {
	addr := fmt.Sprintf(":%s", os.Getenv("PORT"))
	if addr == ":" {
		addr = ":8080" // <1>
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/translate/hello", rest.TranslateHandler)
	mux.HandleFunc("/health", handlers.HealthCheck)
	mux.HandleFunc("/info", handlers.Info)

	log.Printf("listening on %s\n", addr) // <4>

	log.Fatal(http.ListenAndServe(addr, mux)) // <5>
}
