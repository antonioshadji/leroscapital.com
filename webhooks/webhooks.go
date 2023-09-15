package webhooks

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func Handler(w http.ResponseWriter, r *http.Request) {
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
