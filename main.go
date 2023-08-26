package main

import (
	"context"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	secretmanagerpb "google.golang.org/genproto/googleapis/cloud/secretmanager/v1"

	"github.com/antonioshadji/leroscapital.com/treasury"
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
	apiKey = os.Getenv("API_KEY")
	if apiKey == "" {
		// use secret manager
		name := "projects/584752879666/secrets/MAPAPI/versions/2"
		apiKey = accessSecretVersion(name)
	}
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

// accessSecretVersion accesses the payload for the given secret version if one
// exists. The version can be a version number as a string (e.g. "5") or an
// alias (e.g. "latest").
func accessSecretVersion(name string) string {
	// name := "projects/584752879666/secrets/MAPAPI/versions/1"
	// name := "projects/my-project/secrets/my-secret/versions/5"
	// name := "projects/my-project/secrets/my-secret/versions/latest"

	// Create the client.
	ctx := context.Background()
	client, err := secretmanager.NewClient(ctx)
	if err != nil {
		log.Print(fmt.Errorf("failed to create secretmanager client: %v", err))
	}
	defer client.Close()

	// Build the request.
	req := &secretmanagerpb.AccessSecretVersionRequest{
		Name: name,
	}

	// Call the API.
	result, err := client.AccessSecretVersion(ctx, req)
	if err != nil {
		log.Print(fmt.Errorf("failed to access secret version: %v", err))
	}

	return string(result.Payload.Data)
}

func main() {

	http.HandleFunc("/acquisitions/", createHandler("acquisitions"))
	http.HandleFunc("/consulting/", createHandler("consulting"))
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
