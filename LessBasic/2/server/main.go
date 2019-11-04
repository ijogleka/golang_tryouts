package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	// Mock helper struct to marshal the mock message into a json object
	type testMessage struct {
		Message string `json:"message"`
	}

	// Serving the route `localhost:8080/api/v1/test`
	http.HandleFunc("/api/v1/test", func(w http.ResponseWriter, r *http.Request) {
		// Create a mock json response for the test
		b, err := json.Marshal(testMessage{"A test message."})
		if err != nil {
			http.Error(w, "Internal Server Error", 500)
			return
		}

		// Set and send the response to the http.ResponseWriter
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, string(b))
	})

	// This binary will start listening to any messages that it receives
	// on the port 8080
	http.ListenAndServe(":8080", nil)
}
