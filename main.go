package main

import (
  "html/template"
  "net/http"
  "google.golang.org/appengine"
//  "google.golang.org/appengine/datastore"
//  "google.golang.org/appengine/log"
)

type Todo struct {
  Title string
  Done  bool
}

type TodoPageData struct {
  PageTitle string
  Todos     []Todo
}

func Server(w http.ResponseWriter, r *http.Request) {

  data := TodoPageData{
    PageTitle: "Leros Capital",
    Todos: []Todo{
      {Title: "Task 1", Done: false},
      {Title: "Task 2", Done: true},
      {Title: "Task 3", Done: true},
    },
  }

  tmpl := template.Must(template.ParseGlob("templates/*"))
  tmpl.ExecuteTemplate(w, "home", data)
}

func main() {

  http.HandleFunc("/", Server)
  appengine.Main()
  // log.Fatal(http.ListenAndServe(":8080", nil))
}
