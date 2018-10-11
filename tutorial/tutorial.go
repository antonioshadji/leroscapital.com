package tutorial

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"time"

	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/log"
)

type templateParams struct {
	Notice  string
	Name    string
	Message string
	Posts   []Post
}

type Post struct {
	Author  string
	Message string
	Posted  time.Time
}

var tmpl *template.Template

func init() {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(dir)
	tmpl = template.Must(template.ParseFiles("templates/test.tmpl"))
}

func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Sprintf(r.URL.Path)
	if r.URL.Path != "/tutorial" {
		http.Redirect(w, r, "/tutorial", http.StatusFound)
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
			log.Errorf(ctx, "execute template: %v", err)
		}
		return
	}

	if r.Method == "GET" {
		err := tmpl.ExecuteTemplate(w, "test", params)
		if err != nil {
			log.Errorf(ctx, "execute template: %v", err)
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
			log.Errorf(ctx, "execute template: %v", err)
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
			log.Errorf(ctx, "execute template: %v", err)
		}
		return
	}

	// Prepend the post that was just added.
	params.Posts = append([]Post{post}, params.Posts...)

	params.Notice = fmt.Sprintf("Thank you for your submission, %s!", post.Author)
	err := tmpl.ExecuteTemplate(w, "test", params)
	if err != nil {
		log.Errorf(ctx, "execute template: %v", err)
	}
}
