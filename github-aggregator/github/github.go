package github

import (
	"net/http"
	"fmt"
	"errors"
	"io/ioutil"
	"encoding/json"
	"os/exec"
)

type Repository struct {
	URL string `json:"ssh_url"`
}

type RepositoriesManager struct {}

func (r RepositoriesManager) List(user string) ([]Repository, error) {
	resp, err := http.Get(fmt.Sprintf("https://api.github.com/users/%v/repos", user))
	if err != nil {
		return nil, errors.New(fmt.Sprintf("user %v was not found", user))
	}
	body, _ := ioutil.ReadAll(resp.Body)
	repos := []Repository{}
	json.Unmarshal(body, &repos)
	return repos, nil
}

func (r RepositoriesManager) Clone(repo Repository) error {
	dir := "D:\\Users\\aubauman\\Desktop"
	cmd := exec.Command("git", "clone", repo.URL)
	cmd.Dir = dir
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}
