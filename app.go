package main

import (
	"fmt"
	"github.com/antonioshadji/leroscapital.com/treasury"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"
)

var (
	tmpl = template.Must(template.ParseGlob("templates/*"))
)

type PageDetails struct {
	PageTitle string
	Posted    time.Time
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	data := PageDetails{
		PageTitle: "Leros Capital",
		Posted:    time.Now(),
	}

	err := tmpl.ExecuteTemplate(w, "home", data)
	if err != nil {
		log.Printf("Failed to ExecuteTemplate: %v", err)
	}
}

func cbHandler(w http.ResponseWriter, r *http.Request) {
	data := PageDetails{
		PageTitle: "Leros Capital - logged in",
	}

	err := tmpl.ExecuteTemplate(w, "home", data)
	if err != nil {
		log.Printf("Failed to ExecuteTemplate: %v", err)
	}
}

func main() {
	http.HandleFunc("/treasury", treasury.Handler)
	http.HandleFunc("/oath2callback", cbHandler)
	http.HandleFunc("/", homeHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
