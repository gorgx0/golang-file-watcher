package main

import (
	"fmt"
	"github.com/go-git/go-git/v5"
	. "github.com/go-git/go-git/v5/_examples"
	"github.com/go-git/go-git/v5/plumbing/object"
	"time"
)

func main() {
	open, err := git.PlainOpen("/tmp/aaa")
	CheckIfError(err)
	worktree, err := open.Worktree()
	CheckIfError(err)
	_, err = worktree.Add("sss")
	CheckIfError(err)

	status, err := worktree.Status()
	CheckIfError(err)
	fmt.Println(status)

	commit, err := worktree.Commit("file commit", &git.CommitOptions{
		Author: &object.Signature{
			Name:  "Me MK",
			Email: "me@git.com",
			When:  time.Now(),
		},
	})
	CheckIfError(err)
	commitObject, err := open.CommitObject(commit)
	CheckIfError(err)
	fmt.Println(commitObject)
}
