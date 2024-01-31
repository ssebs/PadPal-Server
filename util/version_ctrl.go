// version_ctrl.go - manages version control stuff
package util

import (
	"errors"
	"fmt"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/config"
)

// InitGitDir will check if a git repo for d exists
// It will git init if not.
func InitGitDir(d string) error {
	foundGitRepo := false
	// Get files
	files, err := GetFilenamesInDir(d)
	if err != nil {
		return err
	}
	// check if there's a .git/ folder
	for _, filename := range files {
		if filename == ".git" {
			fmt.Println("found git repo")
			foundGitRepo = true
		}
	}

	// Try to git init, if there's a repo already then open it
	// If there's another error, return. If not, create "main" branch
	repo, err := git.PlainInit(d+"/.git/", foundGitRepo)
	if errors.Is(err, git.ErrRepositoryAlreadyExists) {
		// If the repo exists, open it
		repo, err = git.PlainOpen(d + "/.git/")
		if err != nil {
			return err
		}
	} else if err == nil {
		fmt.Println("Creating main branch on first run")
		// Create main branch
		if err = repo.CreateBranch(&config.Branch{Name: "main"}); err != nil {
			return err
		}
	} else {
		return err
	}

	c, err := repo.Config()
	if err != nil {
		return err
	}

	// b, err := repo.Branch("master")
	// if err != nil {
	// 	return err
	// }

	fmt.Println("config", c.Branches)
	// fmt.Println("branch", b)

	return nil
}
