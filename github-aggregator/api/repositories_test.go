package api

import (
	"net/http/httptest"
	"net/http"
	"sync/atomic"
	"testing"

	"github.com/gorilla/mux"
	"github.com/aubm/introducing-golang/github-aggregator/github"
)

type mockListCloner struct {
	LastListArg string
	NbCloneCalls int32
}

func (lc *mockListCloner) List(user string) ([]github.Repository, error) {
	lc.LastListArg = user
	return []github.Repository{ {URL: "foo"}, {URL: "bar"} }, nil
}

func (lc *mockListCloner) Clone(github.Repository) error {
	lc.NbCloneCalls = atomic.AddInt32(&lc.NbCloneCalls, 1)
	return nil
}

func TestCloneUserRepositories(t *testing.T) {
	// Given
	listCloner := &mockListCloner{}
	handler := RepositoriesHandler{ListCloner: listCloner}
	w := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", "/users/aubm/repos/clone", nil)

	// When
	m := mux.NewRouter()
	m.HandleFunc("/users/{user}/repos/clone", handler.CloneUserRepositories)
	m.ServeHTTP(w, r)

	// Then
	if listCloner.LastListArg != "aubm" {
		t.Errorf("last list arg is %v, expected aubm", listCloner.LastListArg)
	}
	if listCloner.NbCloneCalls != 2 {
		t.Errorf("nb clone calls is %v, expected 2", listCloner.NbCloneCalls)
	}
}
