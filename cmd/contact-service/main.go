package main

import (
	"log"
	"net/http"
	"os"

	"github.com/skill-wanderer/contact-BE/internal/contact"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	svc := contact.NewService()
	mux := http.NewServeMux()

	mux.HandleFunc("GET /healthz", func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"status":"ok"}`))
	})

	mux.HandleFunc("GET /contacts", svc.ListContacts)
	mux.HandleFunc("POST /contacts", svc.CreateContact)

	addr := ":" + port
	log.Printf("contact-service listening on %s", addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatal(err)
	}
}
