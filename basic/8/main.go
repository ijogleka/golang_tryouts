package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"

	"github.com/braintree/manners"
)

func main() {
	// Create a handler (with a ServeHTTP interface attached)
	handler := newHandler()

	// Create a channel, to transfer any OS signals between
	// the main thread and the listener application thread
	ch := make(chan os.Signal)

	// Send a notification over the channel on receiving either of the
	// `Interrupt` or `Kill` signals from the OS (keyboard)
	signal.Notify(ch, os.Interrupt, os.Kill)

	// Start the listener thread
	go listenForShutdown(ch)

	// Start the server in a mannered way :)
	manners.ListenAndServe(":8080", handler)
}

// Just a factory method to create a handler instance
func newHandler() *handler {
	return &handler{}
}

// Declare a request handler struct
type handler struct{}

// Attach a interface to the `handler` struct to `ServeHTTP`
func (h *handler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	// Create a query object
	query := req.URL.Query()

	// Get the Name from the request payload
	name := query.Get("name")

	// If name could not be found, use `Indrajeet`
	if name == "" {
		name = "Indrajeet"
	}

	// Send back a hello text to the client via the ResponseWriter
	fmt.Fprint(res, "Hello, my name is ", name)
}

// Start a application thread to simply listen to Signals from the OS
// Things like a `CTRL+C` key command will trigger a signal from the OS
// Receive any signals on the `ch` channel
func listenForShutdown(ch <-chan os.Signal) {
	<-ch

	// Send a message to shutdown to the server
	// Basically, close the server in a mannered way :)
	manners.Close()
}
