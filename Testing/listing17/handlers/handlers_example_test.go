// Sample test to show how to write a basic example
package handlers_test

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
)

//ExampleSendJSON provides a basic example
//NB an example is always based on an exported function or method
func ExampleSendJSON() {
	r, _ := http.NewRequest("GET", "/sendjson", nil)
	rw := httptest.NewRecorder()
	http.DefaultServeMux.ServeHTTP(rw, r)

	var u struct {
		Name  string
		Email string
	}

	if err := json.NewDecoder(rw.Body).Decode(&u); err != nil {
		log.Println("ERROR: ", err)
	}

	// Use fmt to write to stdout to check the output.
	fmt.Println(u)
	// Output:
	// {Robel robel@golang.com}

}
