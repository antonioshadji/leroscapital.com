package main

import (
	"fmt"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
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

type Post struct {
	Author  string
	Message string
	Posted  time.Time
}

type templateParams struct {
	Notice  string
	Name    string
	Message string
	Posts   []Post
}

func homeHandler(w http.ResponseWriter, r *http.Request) {

	data := PageDetails{
		PageTitle: "Leros Capital",
	}

	err := tmpl.ExecuteTemplate(w, "home", data)
	if err != nil {

	}
}

func testHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Sprintf(r.URL.Path)
	if r.URL.Path != "/test" {
		http.Redirect(w, r, "/test", http.StatusFound)
		return
	}
	ctx := appengine.NewContext(r)
	params := templateParams{}
	q := datastore.NewQuery("Post").Order("-Posted").Limit(20)

	if _, err := q.GetAll(ctx, &params.Posts); err != nil {
		log.Errorf(ctx, "Getting posts: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		params.Notice = "Couldn't get latest posts. Refresh?"
		err := tmpl.ExecuteTemplate(w, "test", params)
		if err != nil {
			log.Errorf(ctx, "execute template: ", err)
		}
		return
	}

	if r.Method == "GET" {
		err := tmpl.ExecuteTemplate(w, "test", params)
		if err != nil {
			log.Errorf(ctx, "execute template: ", err)
		}
		return
	}

	// It's a POST request, so handle the form submission.
	post := Post{
		Author:  r.FormValue("name"),
		Message: r.FormValue("message"),
		Posted:  time.Now(),
	}

	if post.Author == "" {
		post.Author = "Anonymous Gopher"
	}
	params.Name = post.Author

	if post.Message == "" {
		w.WriteHeader(http.StatusBadRequest)
		params.Notice = "No message provided"
		err := tmpl.ExecuteTemplate(w, "test", params)
		if err != nil {
			log.Errorf(ctx, "execute template: ", err)
		}
		return
	}

	key := datastore.NewIncompleteKey(ctx, "Post", nil)

	if _, err := datastore.Put(ctx, key, &post); err != nil {
		log.Errorf(ctx, "datastore.Put: %v", err)

		w.WriteHeader(http.StatusInternalServerError)
		params.Notice = "Couldn't add new post. Try again?"
		params.Message = post.Message // Preserve their message so they can try again.
		err := tmpl.ExecuteTemplate(w, "test", params)
		if err != nil {
			log.Errorf(ctx, "execute template: ", err)
		}
		return
	}

	// Prepend the post that was just added.
	params.Posts = append([]Post{post}, params.Posts...)

	params.Notice = fmt.Sprintf("Thank you for your submission, %s!", post.Author)
	err := tmpl.ExecuteTemplate(w, "test", params)
	if err != nil {
		log.Errorf(ctx, "execute template: ", err)
	}
}

func main() {
	http.HandleFunc("/test", testHandler)
	http.HandleFunc("/", homeHandler)
	appengine.Main()
}
