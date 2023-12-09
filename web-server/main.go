package main

import (
	"log"
	"net/http"
	"net/http/httputil"
)

const serverPort = ":9001"
const responseHeaderTestKey = "Test-Header"
const responseHeaderTestValue = "Test-Header-Value"
const responseMessage = "Endpoint successfully called\n"

func main() {
	// Start the server
	startServer()
}

func startServer() {
	log.Println("Starting server on port", serverPort)

	handler := func(rw http.ResponseWriter, r *http.Request) {
		handleRequest(rw, r)
	}

	err := http.ListenAndServe(serverPort, http.HandlerFunc(handler))
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func handleRequest(rw http.ResponseWriter, r *http.Request) {
	// Dump the request data for logging
	reqData, err := httputil.DumpRequest(r, true)
	if err != nil {
		log.Printf("Error dumping request: %v", err)
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte("Internal Server Error\n"))
		return
	}

	log.Printf("Received request at endpoint: %s", string(reqData))

	rw.Header().Add(responseHeaderTestKey, responseHeaderTestValue)
	rw.WriteHeader(http.StatusAccepted)
	rw.Write([]byte(responseMessage))
	log.Printf("Request served with header %s: %s and message: %s", responseHeaderTestKey, responseHeaderTestValue, responseMessage)
}
