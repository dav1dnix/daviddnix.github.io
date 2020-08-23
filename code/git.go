package main

import (
	"github.com/go-git/go-git/v5"
	"log"
	"fmt"
	"flag"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
	"github.com/go-git/go-git/v5/storage/memory"
	"github.com/go-git/go-git/v5/config"
	"github.com/dps910/dps910.github.io/code/functions"
)

var (
	repository = flag.String("repository", "", "Git repository")
	username = flag.String("username", "", "Git Username")
	password = flag.String("password", "", "Git password")
	auth = &http.BasicAuth {
		Username: "dos819",
		Password: "thisismytokenlol/s",
	}
)

func main() {
	flag.Parse()

	// First, a git repository needs to be initialized if it does not already exist.
	// Storage will be a storage base on memory.
	r, _ := git.Init(memory.NewStorage(), nil)

	// Create remote with name of remote and URL of repository
	_, err := r.CreateRemote(&config.RemoteConfig{
		Name: "website",
		URLs: []string{"https://github.com/" + *repository},
	})
	if err != nil {
		log.Fatal(err)
	}

	// Get a list of all remotes
	list, err := r.Remotes()
	if err != nil {
		log.Fatal(err)
	}

	for _, r := range list {
		fmt.Println(r)
	}

	// Open git repository, if path doesn't contain valid repository ErrRepositoryNotExists is returned.
	git.PlainOpen("dps910/dps910.github.io")

	// Push code to remote repository.
	functions.Info("Git Push")
	err = r.Push(&git.PushOptions{
		RemoteName: "website",
		Auth: auth,
	})
	if err != nil {
		log.Fatal(err)
	}
}
