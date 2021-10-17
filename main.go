package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/antonioshadji/leroscapital.com/treasury"
)

var tmpl = template.Must(template.ParseGlob("templates/*"))

// PageDetails ...
type PageDetails struct {
	PageTitle  string
	PageHeader string
	Posted     time.Time
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	data := PageDetails{
		PageTitle:  "Leros Capital",
		PageHeader: "Leros Capital",
		Posted:     time.Now(),
	}

	err := tmpl.ExecuteTemplate(w, "home", data)
	if err != nil {
		log.Printf("Failed to ExecuteTemplate: %v", err)
	}
}

func mapHandler(w http.ResponseWriter, r *http.Request) {
	err := tmpl.ExecuteTemplate(w, "map", PageDetails{})
	if err != nil {
		log.Printf("Failed to ExecuteTemplate: %v", err)
	}
}

func cbHandler(w http.ResponseWriter, r *http.Request) {
	data := PageDetails{
		PageTitle:  "Leros Capital ::",
		PageHeader: "Leros Capital",
		Posted:     time.Now(),
	}

	err := tmpl.ExecuteTemplate(w, "home", data)
	if err != nil {
		log.Printf("Failed to ExecuteTemplate: %v", err)
	}
}

func webhookHandler(w http.ResponseWriter, r *http.Request) {
	log.Print(formatRequest(r))
}

// from: https://medium.com/doing-things-right/pretty-printing-http-requests-in-golang-a918d5aaa000
// formatRequest generates ascii representation of a request
func formatRequest(r *http.Request) string {
	// Create return string
	var request []string
	// Add the request string
	url := fmt.Sprintf("%v %v %v", r.Method, r.URL, r.Proto)
	request = append(request, url)
	// Add the host
	request = append(request, fmt.Sprintf("Host: %v", r.Host))
	// Loop through headers
	for name, headers := range r.Header {
		name = strings.ToLower(name)
		for _, h := range headers {
			request = append(request, fmt.Sprintf("%v: %v", name, h))
		}
	}

	// If this is a POST, add post data
	if r.Method == "POST" {
		body, err := ioutil.ReadAll(r.Body)

		// https://blog.golang.org/json-and-go
		var f interface{}
		err = json.Unmarshal(body, &f)
		if err != nil {
			log.Println("json fail")
			request = append(request, "json fail")
		} else {
			request = append(request, indent(f))
		}

	}
	request = append(request, "=================================================")
	// Return the request as a string
	return strings.Join(request, "\n")
}

func indent(v interface{}) string {
	b, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return fmt.Sprintf("%#v", v)
	}
	return string(b)
}

func main() {
	http.HandleFunc("/treasury/", treasury.Handler)
	http.HandleFunc("/oath2callback/", cbHandler)
	http.HandleFunc("/webhook/", webhookHandler)
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
