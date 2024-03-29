package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	if req.Method != "GET" {
		http.Error(w, "method is not supported", http.StatusMethodNotAllowed)
		return
	}

	fmt.Fprintf(w, "hello!")
}

func formHandler(w http.ResponseWriter, req *http.Request) {
	if err := req.ParseForm(); err != nil {
		fmt.Fprintf(w, "Parse form err: %v", err)
		return
	}

	fmt.Fprintf(w, "post request successful\n")
	name := req.FormValue("name")
	address := req.FormValue("address")
	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Address = %s\n", address)
}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
