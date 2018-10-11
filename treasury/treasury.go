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

// headers returned
//  map[
// 	X-Content-Type-Options:[nosniff]
// 	Transfer-Encoding:[chunked]
// 	Set-Cookie:[BIGipServerpl_www.treasurydirect.gov_443=3204804106.47873.0000;path=/; Httponly; Secure TS01598982=019e2ba2e91750a71b2f37c19d9059c6b6caa5b9a9677379d8c7d1c8355573a91fde5ccbe33210d44123d567968738d32d7bc8aff026e3d7bd460fbaa436e7fa806031fc2b; Path=/]
// 	Strict-Transport-Security:[max-age=31536000; includeSubDomains]
// 	Cache-Control:[no-store]
// 	X-Frame-Options:[SAMEORIGIN]
// 	Surrogate-Control:[content="ESI/1.0",no-store]
// 	Date:[Thu, 11 Oct 2018 16:55:31 GMT]
// 	Content-Type:[application/json;charset=UTF-8]
// 	Content-Language:[en-US]
// 	]
