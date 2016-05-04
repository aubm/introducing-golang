package main

import (
	"net/http"
	"github.com/aubm/introducing-golang/github-aggregator/api"
	"log"
	"fmt"
)

func main() {
	http.HandleFunc("/hello", api.Hello)
	fmt.Println("Application started on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

