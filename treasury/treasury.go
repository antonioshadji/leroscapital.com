package treasury

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"google.golang.org/appengine"
	"google.golang.org/appengine/urlfetch"
)

func init() {}

func Handler(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	client := urlfetch.Client(ctx)
	resp, err := client.Get("https://www.treasurydirect.gov/NP_WS/debt/current?format=json")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "HTTP GET returned status: %v\n", resp.Status)
	fmt.Fprintf(w, "HTTP GET returned headers:\n %v\n", resp.Header)
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Fprintf(w, "HTTP GET returned body:\n %v\n", string(body))
}
