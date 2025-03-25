package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if (r.URL.Path) != "/hello" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
        http.Error(w, "Method is not supported.", http.StatusNotFound)
        return
    }

	fmt.Fprintf(w, "Hello!")
}

func main() {
    // Create a new ServeMux (router)
    router := http.NewServeMux()

    // Register routes
    router.HandleFunc("/hello", helloHandler)

    // Create server
    server := &http.Server{
        Addr:    ":8080",
        Handler: router,
    }

    fmt.Printf("Starting server at port 8080\n")
    if err := server.ListenAndServe(); err != nil {
        log.Fatal(err)
    }
}