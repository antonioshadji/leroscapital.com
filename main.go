package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/antonioshadji/leroscapital.com/secrets"
	"github.com/antonioshadji/leroscapital.com/treasury"
	"github.com/antonioshadji/leroscapital.com/webhooks"
)

var tmpl = template.Must(template.ParseGlob("templates/*"))
var apiKey string

// PageDetails ...
type PageDetails struct {
	PageTitle  string
	PageHeader string
	Posted     time.Time
	APIKey     string
	Year       int
}

var data = PageDetails{
	PageTitle:  "Leros Capital LLC",
	PageHeader: "Leros Capital",
	Posted:     time.Now(),
	Year:       time.Now().Year(),
}

func init() {
	name := "projects/584752879666/secrets/MAPAPI/versions/2"
	apiKey = secrets.AccessSecretVersion(name)
}

func createHandler(path string) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		data.PageTitle = "Leros Capital :: " + strings.Title(path)

		err := tmpl.ExecuteTemplate(w, path, data)
		if err != nil {
			log.Printf("Failed to ExecuteTemplate: %v", err)
		}
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	data.PageTitle = "Leros Capital LLC"

	err := tmpl.ExecuteTemplate(w, "home", data)
	if err != nil {
		log.Printf("Failed to ExecuteTemplate: %v", err)
	}
}

func mapHandler(w http.ResponseWriter, r *http.Request) {
	data.APIKey = apiKey

	err := tmpl.ExecuteTemplate(w, "map", data)
	if err != nil {
		log.Printf("Failed to ExecuteTemplate: %v", err)
	}
}

func cbHandler(w http.ResponseWriter, r *http.Request) {
	data.PageTitle = "Leros Capital LLC :: Callback Handler"

	err := tmpl.ExecuteTemplate(w, "home", data)
	if err != nil {
		log.Printf("Failed to ExecuteTemplate: %v", err)
	}
}

func main() {

	http.HandleFunc("/acquisitions/", createHandler("acquisitions"))
	http.HandleFunc("/capabilities/", createHandler("capabilities"))
	http.HandleFunc("/consulting/", createHandler("consulting"))
	http.HandleFunc("/treasury/", treasury.Handler)
	http.HandleFunc("/oauth2callback/", cbHandler)
	http.HandleFunc("/webhooks/", webhooks.Handler)
	http.HandleFunc("/map", mapHandler)
	http.HandleFunc("/", homeHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
