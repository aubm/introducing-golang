package api

import (
	"log"
	"net/http"

	"github.com/aubm/present-talks/introducing-golang/github-aggregator/github"
	"github.com/gorilla/mux"
)

type listCloner interface {
	List(user string) ([]github.Repository, error)
	Clone(github.Repository) error
}

type RepositoriesHandler struct {
	ListCloner listCloner
}

func (h RepositoriesHandler) CloneUserRepositories(w http.ResponseWriter, r *http.Request) {
	requestVars := mux.Vars(r)
	repos, err := h.ListCloner.List(requestVars["user"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	clonesDone := make(chan (bool), len(repos))
	for _, repo := range repos {
		go func(repo github.Repository) {
			if err := h.ListCloner.Clone(repo); err == nil {
				log.Printf("repo %v successfully cloned", repo)
			} else {
				log.Printf("error when cloning repo %v: %v", repo, err)
			}
			clonesDone <- true
		}(repo)
	}
	for range repos {
		<-clonesDone
	}
	log.Println("all clones done")

	w.WriteHeader(http.StatusTeapot)
}
