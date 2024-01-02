# Simple Go Web Server

A basic Go web server handling static files and two endpoints: `/hello` and `/form`.

## Usage

- Clone/download the repository.
- Ensure Go is installed.
- Navigate to the project directory.
- Run `go run main.go` to start the server.
- Access `http://localhost:8080` in a browser to view served files.

## Endpoints

- `/` : Responds with "Static website!" content.
- `/hello`: Responds with "hello!" to a GET request.
- `/form`: Handles POST requests with `name` and `address` parameters.

## Dependencies

Uses standard Go `net/http` package.

