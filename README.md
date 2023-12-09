# Go Reverse Proxy Project

This project contains an implementation of a Reverse Proxy using the Go (Golang) programming language. Plus, it creates a basic HTTP server that the Reverse Proxy interacts with.

## About The Project

The project is composed of two main parts: the Reverse Proxy and a simple web server with which the proxy sends requests.

#### Reverse Proxy

The Proxy code follows the reverse proxy design pattern, which takes client requests and forwards them to the appropriate origin server. It then collects the response from the server and returns it to the client.

The code for the proxy is located in the file `main.go` in the root folder of the project.

#### Web Server

We have a simple web server that the Reverse Proxy interacts with. This program listens on port 9001 and simply responds to all HTTP requests with a 200 status code and a message.

The code for the web server is located in a separate `main.go` file.

## Technical Details

This project has been built using Go version 1.20.6. It touches upon several aspects of the Go language, including handling HTTP, error handling, and reading flags from the command line.

### Reverse Proxy Code

The Reverse Proxy starts by starting an HTTP server listening on port 9000. Any request that comes into this server is handled by the anonymous function inside `http.HandleFunc()`.

This function handles the request by creating a new HTTP request with the same method (GET, POST, etc.) and body as the original request. It then sets the headers of the new request to be the same as those from the original request.

Finally, the new request is sent to the origin server that was provided as a flag on the command line. The response from the origin server is then returned to the client who made the original request.

### Web Server Code

The web server is a simple Go program that upon receiving an HTTP request, returns a response with a 200 status code and a message.

## How To Use

To start the Reverse Proxy, run the command in root folder:

```
go run main.go -e http://localhost:9001
```

To start the Web Server, navigate to the web-server folder and run the command:
```
go run main.go
```

## Contributions
Contributions are welcome. Please open an 'Issue' or 'Pull Request' for your contributions.

## License
This project is under the MIT license.