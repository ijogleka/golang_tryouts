package main

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
)

// hasTimedOut helps with checking if a input error object represents
// a http request timeout
func hasTimedOut(err error) bool {
	switch err := err.(type) {
	case *url.Error:
		if err, ok := err.Err.(net.Error); ok && err.Timeout() {
			return true
		}
	case net.Error:
		if err.Timeout() {
			return true
		}
	case *net.OpError:
		if err.Timeout() {
			return true
		}
	}

	errTxt := "use of closed network connection"
	if err != nil && strings.Contains(err.Error(), errTxt) {
		return true
	}
	return false
}

// This is a recursive function to download full (or parts) of remote file
// onto the local filesystem
//
// @param location [string] remote url to download the file from
// @param file [*os.File] pointer to a local file to download remote file to
// @param retries [int64] number of allowed download retries, in case of timeouts
//
// @return error [error object] or nil
func download(location string, file *os.File, retries int64) error {
	// Create a http GET request
	req, err := http.NewRequest("GET", location, nil)
	if err != nil {
		return err
	}

	// Check the local file size to see if some parts of the file have already
	// been downloaded
	fi, err := file.Stat()
	if err != nil {
		return err
	}
	current := fi.Size()
	if current > 0 {
		start := strconv.FormatInt(current, 10)
		// http request header to indicate which byte offset to start reading
		// the remote file from.
		// - When no part of the file has been downloaded yet, this will not be set
		req.Header.Set("Range", "bytes="+start+"-")
	}

	// Create a http client (with a timeout configuration of 5 minutes)
	cc := &http.Client{Timeout: 5 * time.Minute}

	// Perform the http GET request
	res, err := cc.Do(req)
	if err != nil && hasTimedOut(err) {
		// In case the request timed out while trying to download the file,
		// retry downloading the remaining bytes of the file.
		if retries > 0 {
			return download(location, file, retries-1)
		}
		// If retry attempts have reached our limit and the file has still not been
		// completely downloaded, return a error object
		return err
	} else if err != nil {
		// If the http request failed, but the reason was not a timeout, return the
		// error object anyway
		return err
	}

	// Check if the http request was successful
	if res.StatusCode < 200 || res.StatusCode >= 300 {
		errFmt := "Unsuccess HTTP request. Status: %s"
		return fmt.Errorf(errFmt, res.Status)
	}

	// If the server dows not support downloading the file in parts,
	// set the number of allowed retries to 0
	if res.Header.Get("Accept-Ranges") != "bytes" {
		retries = 0
	}

	// For a successful response, copy the body of the response (i.e. some part of the remote file)
	// to the local file
	_, err = io.Copy(file, res.Body)
	if err != nil && hasTimedOut(err) {
		// If the copy fails due to a timeout, retry (until allowed)
		if retries > 0 {
			return download(location, file, retries-1)
		}
		return err
	} else if err != nil {
		return err
	}

	return nil
}

func main() {
	// Create a zip file on the local filesystem
	file, err := os.Create("file.zip")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	// Download a zip file from a remote url location
	location := "https://example.com/file.zip"
	err = download(location, file, 100)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Check the file statistics, in this example, the size
	fi, err := file.Stat()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Got it with %v bytes downloaded", fi.Size())
}
