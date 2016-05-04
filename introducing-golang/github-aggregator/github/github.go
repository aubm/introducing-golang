package github

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os/exec"
)

type Config struct {
	Dir string
}

type Repository struct {
	URL string `json:"ssh_url"`
}

type RepositoriesManager struct {
	Config
}

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
	cmd := exec.Command("git", "clone", repo.URL)
	cmd.Dir = r.Dir
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}
