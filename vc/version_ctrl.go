// version_ctrl.go - manages version control stuff
package vc

import (
	"errors"
	"fmt"
	"time"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
)

// InitGitDir will check if a git repo for d exists
// It will git init if not.
func InitGitDir(d string) (*git.Repository, error) {
	// Try to git init, if there's a repo already then open it
	// If there's another error, return. If not, create "main" branch
	repo, err := git.PlainInit(d, false)
	if errors.Is(err, git.ErrRepositoryAlreadyExists) {
		// If the repo exists, open it
		repo, err = git.PlainOpen(d)
		if err != nil {
			return repo, err
		}
	} else if err != nil {
		return nil, err
	}

	return repo, nil
}

// AddCommitFile will git add & git commit a file to the repo
func AddCommitFile(file, author string, repo *git.Repository) error {
	// Create staging area / worktree
	wt, err := repo.Worktree()
	if err != nil {
		return fmt.Errorf("failed to create staging area within repo, err: %s", err.Error())
	}
	// Add the file
	_, err = wt.Add(file)
	if err != nil {
		return fmt.Errorf("failed to add file, err: %s", err.Error())
	}
	// Commit the file
	commitMsg := fmt.Sprintf("Commiting %s", file[len("active/"):])
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

	repo.CommitObject(commit)
	return nil
}

// AddCommitDir will git add & git commit a file to the repo
func AddCommitDir(dir, author string, repo *git.Repository) error {
	// Create staging area / worktree
	wt, err := repo.Worktree()
	if err != nil {
		return fmt.Errorf("failed to create staging area within repo, err: %s", err.Error())
	}
	// Add the file
	_, err = wt.Add(dir)
	if err != nil {
		return fmt.Errorf("failed to add file, err: %s", err.Error())
	}
	// Commit the file
	commitMsg := fmt.Sprintf("Commiting %s", dir)
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

	repo.CommitObject(commit)
	return nil
}
