package api

import "net/http"

type ReposHandlers struct{}

func (h ReposHandlers) CloneRepos(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, Go!"))
}
