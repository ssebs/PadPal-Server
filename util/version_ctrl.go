// version_ctrl.go - manages version control stuff
package util

import (
	"fmt"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/storage/memory"
)

func InitGitDir(d string) error {
	files, err := GetFilenamesInDir(d)
	if err != nil {
		return err
	}
	// check if there's a .git/ folder
	for _, filename := range files {
		fmt.Println(filename)
		// If we found it, great we can leave
		if filename == ".git" {
			fmt.Println("found")
			return nil
		}
	}
	// if not, create one
	repo, err := git.Init(memory.NewStorage(), nil)
	// repo.

	return nil
}
