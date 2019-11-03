package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// This function demos a simple usage of the default http client from the golang library
func basicHTTPGetter() {
	// This is a default http client provided in the golang standard library
	// NOTE: The default client has configuration and functionality to handle things
	//       like HTTP redirects, cookies, and timeouts
	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	// returned response contains headers, encoding, body etc
	// returned `err` object can be `nil` or a `error` object
	fmt.Println("ErrorCode for the Get request:", err)

	// Read `Body` of the http response into a local variable `b`
	b, _ := ioutil.ReadAll(res.Body)
	res.Body.Close() // Close the response stream?

	// NOTE: Using `Println` here results in printing of ascii values
	fmt.Printf("%s", b)
}

// This function demos the usage of a customized http client (configuration customized)
func basicHTTPClient() {
	// Create a http GET Request
	req, _ := http.NewRequest("GET", "https://jsonplaceholder.typicode.com/todos/1", nil)

	// Add any custom headers
	req.Header.Add("If-None-Match", `W/"wyzzy"`)

	// Create a http client with configurations like `redirect policy`,
	// `Request timeout`, `keep-alives`, `compression` etc
	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
	}
	client := &http.Client{
		// CheckRedirect: redirectPolicyFunc,
		Transport: tr,
	}

	// Perform the http request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("HTTP Request error:", err)
	}
	// Close the Body (stream) of the response once read (i.e. end of the scope)
	defer resp.Body.Close()

	// Read the response body into a local variable
	body, _ := ioutil.ReadAll(resp.Body)

	// Print the body data
	fmt.Printf("%s", body)
}

func main() {
	basicHTTPGetter()
	basicHTTPClient()
}

// NOTE: Clients and Transports are safe for concurrent use by multiple goroutines
// and for efficiency should only be created once and re-used.

// NOTE: HTTP Post looks something like
//  resp, err := http.Post("http://example.com/upload", "image/jpeg", &buf)
// OR
//  resp, err := http.PostForm("http://example.com/form",
//      url.Values{"key": {"Value"}, "id": {"123"}})
