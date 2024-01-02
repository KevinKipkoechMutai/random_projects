package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w, r *)  {
	if r.URL.path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	if r.method != "GET" {
		http.Error(w, "method is not supported", http.StatusNotFound)
	}

	fmt.fprintf(w, "hello!")
}

func formHandler()  {
	
}

func main()  {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServer(":8080", nil); err != nil {
		log.Fatal(err)
	}
}