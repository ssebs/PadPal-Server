// version_ctrl.go - manages version control stuff
package vc

import (
	"errors"
	"fmt"
	"time"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
)

type VersionControl struct {
	baseDir string
	repo    *git.Repository
}

// NewVersionControl will check if a git repo for d exists, will git init if not.
func NewVersionControl(d string) (*VersionControl, error) {
	vc := &VersionControl{
		baseDir: d,
	}
	// Try to git init, if there's a repo already then open it
	// If there's another error, return. If not, create "main" branch
	repo, err := git.PlainInit(d, false)
	if errors.Is(err, git.ErrRepositoryAlreadyExists) {
		// If the repo exists, open it
		repo, err = git.PlainOpen(d)
		if err != nil {
			return nil, err
		}
	} else if err != nil {
		return nil, err
	}

	vc.repo = repo
	return vc, nil
}

// AddCommit will git add & git commit a file to the repo
// file can be the local path to the file relative to v.baseDir
// author is the commit author
// isDir is whether you're adding a directory or a file
func (v *VersionControl) AddCommit(file, author string, isDir bool) error {
	commitMsg := ""
	// Create staging area / worktree
	wt, err := v.repo.Worktree()
	if err != nil {
		return fmt.Errorf("failed to create staging area within repo, err: %s", err.Error())
	}
	// Add the file
	_, err = wt.Add(file)
	if err != nil {
		return fmt.Errorf("failed to add file, err: %s", err.Error())
	}

	// Commit the file
	if isDir {
		commitMsg = fmt.Sprintf("Commiting %s", file)
	} else {
		commitMsg = fmt.Sprintf("Commiting %s", file[len("active/"):])
	}

	commit, err := wt.Commit(commitMsg, &git.CommitOptions{
		Author: &object.Signature{
			Name:  author,
			Email: author + "@example.com",
			When:  time.Now().UTC(),
		},
	})
	if err != nil {
		return fmt.Errorf("failed to commit file, err: %s", err.Error())
	}

	v.repo.CommitObject(commit)
	return nil
}
