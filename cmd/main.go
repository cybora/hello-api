package main

import (
	"log"
	"net/http"

	"github.com/cybora/shipping_go/handlers"
	"github.com/cybora/shipping_go/handlers/rest"
	"github.com/cybora/shipping_go/translation"
)

func main() {
	addr := ":8080"

	mux := http.NewServeMux()

	translationService := translation.NewStaticService()
	translationHandler := rest.NewTranslateHandler(translationService)

	mux.HandleFunc("/hello", translationHandler.TranslateHandler)
	mux.HandleFunc("/health", handlers.HealthCheck)

	log.Printf("Listening on %s\n", addr)
	log.Fatal(http.ListenAndServe(addr, mux))
}
