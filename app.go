package main

import (
	"github.com/antonioshadji/leroscapital.com/treasury"
	"github.com/antonioshadji/leroscapital.com/tutorial"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	"html/template"
	"net/http"
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
	ctx := appengine.NewContext(r)
	data := PageDetails{
		PageTitle: "Leros Capital",
		Posted:    time.Now(),
	}

	err := tmpl.ExecuteTemplate(w, "home", data)
	if err != nil {
		log.Errorf(ctx, "Failed to ExecuteTemplate: %v", err)
	}
}

func cbHandler(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	data := PageDetails{
		PageTitle: "Leros Capital - logged in",
	}

	err := tmpl.ExecuteTemplate(w, "home", data)
	if err != nil {
		log.Errorf(ctx, "Failed to ExecuteTemplate: %v", err)
	}
}

func main() {
	http.HandleFunc("/treasury", treasury.Handler)
	http.HandleFunc("/tutorial", tutorial.Handler)
	http.HandleFunc("/oath2callback", cbHandler)
	http.HandleFunc("/", homeHandler)
	appengine.Main()
}
