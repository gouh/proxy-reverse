package main

import (
	"flag"
	"io"
	"log"
	"net/http"
	"net/http/httputil"
	"time"
)

const ServerTimeout = 5 * time.Second
const ServerPort = ":9000"

func main() {
	// Get the endpoint as a flag
	endpoint := processFlags()

	// Start the server
	startServer(endpoint)
}

// processFlags captures endpoint from command line flags
func processFlags() *string {
	endpoint := flag.String("e", "", "")
	flag.Parse()

	if *endpoint == "" {
		log.Fatal("endpoint is required")
	}

	return endpoint
}

// startServer starts the HTTP server at ServerPort
func startServer(endpoint *string) {
	log.Println("starting server")

	handler := func(rw http.ResponseWriter, r *http.Request) {
		forwardRequest(endpoint, rw, r)
	}

	// On every request call this handler
	log.Fatal(http.ListenAndServe(ServerPort, http.HandlerFunc(handler)))
}

// forwardRequest forwards requests to the specified endpoint
func forwardRequest(endpoint *string, rw http.ResponseWriter, r *http.Request) {
	// Create a new request
	req, err := http.NewRequest(r.Method, *endpoint, r.Body)
	if err != nil {
		processErr(rw, err)
		return
	}

	prepareRequest(r, req)

	// Actually forward the request to our endpoint
	resp := doRequest(req, rw)
	if resp == nil {
		return
	}
	defer resp.Body.Close()

	processResponse(resp, rw)
}

// prepareRequest prepares an HTTP request by cloning headers and setting query parameters
func prepareRequest(r *http.Request, req *http.Request) {
	// Set the request headers as whatever was passed by caller
	req.Header = r.Header.Clone()

	// Set the query parameters
	req.URL.RawQuery = r.URL.RawQuery
}

// doRequest sends our request and returns the response
func doRequest(req *http.Request, rw http.ResponseWriter) *http.Response {
	// Create a http client, timeout should be mentioned or it will never timeout.
	client := http.Client{
		Timeout: ServerTimeout,
	}

	resp, err := client.Do(req)
	if err != nil {
		processErr(rw, err)
		return nil
	}

	return resp
}

// processResponse processes the forwarded request's response
func processResponse(resp *http.Response, rw http.ResponseWriter) {
	// Get dump of our response
	respData, err := httputil.DumpResponse(resp, true)
	if err != nil {
		processErr(rw, err)
		return
	}

	log.Println("Forward Request Response", string(respData))

	// Copy the response headers to the actual response. DO THIS BEFORE CALLING WRITEHEADER.
	for k, v := range resp.Header {
		rw.Header()[k] = v
	}

	// Set the status code whatever we got from the response
	rw.WriteHeader(resp.StatusCode)

	// Copy the response body to the actual response
	_, err = io.Copy(rw, resp.Body)
	if err != nil {
		log.Println(err)
		rw.Write([]byte("error"))
	}
}

// processErr logs the error and writes an HTTP Internal Server Error response
func processErr(rw http.ResponseWriter, err error) {
	log.Println(err)
	rw.WriteHeader(http.StatusInternalServerError)
	rw.Write([]byte("error"))
}
