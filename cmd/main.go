package main

import (
	"fmt"
	"os"
	"time"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
	. "github.com/lovelacelee/clsgo/pkg/log"
	. "github.com/lovelacelee/clsgo/pkg/utils"
)

func main() {

	author := &object.Signature{
		Name:  "lovelacelee",
		Email: "lovelaelee@gmail.com",
		When:  time.Now(),
	}
	directory, _ := os.Getwd()

	// Opens an already existing repository.
	r, err := git.PlainOpen(directory)
	ExitIfError(err)
	Info("Worktree: %s\n", directory)
	w, err := r.Worktree()
	ExitIfError(err)

	// Verify the current status of the worktree using the method Status.
	Info("git status --porcelain")
	status, err := w.Status()
	ExitIfError(err)

	fmt.Println(status)

	Info("git add .")
	_, err = w.Add(".")
	ExitIfError(err)

	// Commits the current staging area to the repository.
	// We should provide the object.Signature of Author of the
	// commit Since version 5.0.1, we can omit the Author signature, being read
	// from the git config files.
	Info("git commit")
	commit, err := w.Commit("Self committed", &git.CommitOptions{
		Author: author,
	})
	ExitIfError(err)

	// Prints the current HEAD to verify that all worked well.
	Info("git show -s")
	obj, err := r.CommitObject(commit)
	ExitIfError(err)

	fmt.Println(obj)

	// push using default options
	err = r.Push(&git.PushOptions{})
	ExitIfError(err)
}
