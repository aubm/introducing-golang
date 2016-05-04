package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/aubm/present-talks/introducing-golang/github-aggregator/api"
	"github.com/aubm/present-talks/introducing-golang/github-aggregator/github"
	"github.com/gorilla/mux"
)

var config github.Config

func init() {
	flag.StringVar(&config.Dir, "dir", "", "the target dir for git clone")
	flag.Parse()

	if config.Dir == "" {
		fmt.Println("You must specify a target dir using the -dir command flag")
		os.Exit(1)
	}
}

func main() {
	repositoriesHandler := api.RepositoriesHandler{ListCloner: github.RepositoriesManager{config}}

	router := mux.NewRouter()
	router.HandleFunc("/users/{user}/repos/clone", repositoriesHandler.CloneUserRepositories)

	fmt.Println("Application started on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
