package main

import (
	"fmt"
	"net/http"
	"log"

	"github.com/aubm/introducing-golang/github-aggregator/api"
	"github.com/aubm/introducing-golang/github-aggregator/github"
	"github.com/gorilla/mux"
)

func main() {
	repositoriesHandler := api.RepositoriesHandler{ListCloner: github.RepositoriesManager{}}

	router := mux.NewRouter()
	router.HandleFunc("/users/{user}/repos/clone", repositoriesHandler.CloneUserRepositories)

	fmt.Println("Application started on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
