package treasury

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

var (
	tmpl = template.Must(template.ParseGlob("templates/*"))
)

// PageDetails page data
type PageDetails struct {
	PageTitle  string
	PageHeader string
	Posted     time.Time
}

var data PageDetails

func init() {
	data = PageDetails{
		PageTitle:  "Leros Capital :: TreasuryDirect API",
		PageHeader: "TreasuryDirect API",
		Posted:     time.Now(),
	}
}

// Handler http handler
func Handler(w http.ResponseWriter, r *http.Request) {
	if r.URL.EscapedPath() == "/treasury/" {
		err := tmpl.ExecuteTemplate(w, "treasury", data)
		if err != nil {
			log.Printf("Failed to ExecuteTemplate: %v", err)
		}
		return
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	q := r.URL.Query()

	p := strings.TrimPrefix(r.URL.EscapedPath(), "/treasury")
	var b strings.Builder
	b.WriteString("https://treasurydirect.gov")
	b.WriteString(p)
	b.WriteString("?")
	for key, value := range q {
		_, err := fmt.Fprintf(&b, "%v=%v", key, value[0])
		if err != nil {
			log.Println("query fail")
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	// res, err := http.Get("https://www.treasurydirect.gov/NP_WS/debt/current?format=json")
	res, err := http.Get(b.String())
	if err != nil {
		log.Println("get fail")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	body, err := ioutil.ReadAll(res.Body)

	// https://blog.golang.org/json-and-go
	var f interface{}
	err = json.Unmarshal(body, &f)
	if err != nil {
		log.Println("json fail")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(f)
	if err != nil {
		log.Println("encode fail")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
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
