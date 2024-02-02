// file_provider.go - implement file based CRUDProvider
package data

import (
	"os"
	"path/filepath"
	"sync"

	"github.com/beevik/guid"
	"github.com/go-git/go-git/v5"
	"github.com/ssebs/padpal-server/vc"
)

// FileProvider is a file directory implementation of CRUDProvider
type FileProvider struct {
	fullPath string
	repo     *git.Repository
	notes    map[guid.Guid]*Note
	mutex    sync.RWMutex
}

// NewFileProvider
// make sure to include the last / for the dir!
func NewFileProvider(dir string) (*FileProvider, error) {
	// Check dir exists
	d := filepath.Dir(dir)
	if _, err := os.Stat(d); os.IsNotExist(err) {
		return nil, err
	}

	// Load notes from existing dir + check if we're loading an existing dir or creating new
	_repo, err := initFileProvider(d)
	if err != nil {
		return nil, err
	}

	fp := &FileProvider{
		fullPath: d,
		notes:    make(map[guid.Guid]*Note),
		repo:     _repo,
	}
	// TODO: load notes in ListNotes/LoadNotes
	return fp, nil
}

// initFileProvider will load notes from existing dir +
// Checks if we're loading an existing dir or creating new git dir
func initFileProvider(d string) (*git.Repository, error) {
	// Init git dir + get repo
	repo, err := vc.InitGitDir(d)
	if err != nil {
		return nil, err
	}

	return repo, nil
}

// CREATE //
// Save note to disk
func (p *FileProvider) SaveNote(note *Note) error {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	filename := p.fullPath + "/" + note.GetFilename()

	// save md file to disk
	if err := os.WriteFile(filename, []byte(note.Contents), 0644); err != nil {
		return err
	}
	// git add & commit
	if err := vc.AddCommitFile(filename, note.Author, p.repo); err != nil {
		return err
	}

	return nil
}

// TODO: CopyNote(id guid.Guid) error

// READ //
// List all active notes
func (p *FileProvider) ListNotes(query string) ([]*Note, error) {
	// TODO: MOVE TO LISTNOTES + READFILES + PLAN THIS OUT!!

	// // ...then load notes from "active folder"
	// files, err := util.GetFilesInDir(d)
	// if err != nil {
	// 	return nil, notes, err
	// }
	// // Read contents of file, parse into Note
	// for _, f := range files {
	// 	contents, err := io.ReadAll(f)
	// 	if err != nil {
	// 		return nil, notes, err
	// 	}
	// 	notes[]
	// }
	return nil, nil
}

// Load note from disk by guid ID
func (p *FileProvider) LoadNote(id guid.Guid) (*Note, error) {
	return nil, nil
}

// List all versions of a note
func (p *FileProvider) ListNoteVersions(id guid.Guid) ([]int, error) {
	return nil, nil
}

// Load note from disk by guid ID + version
func (p *FileProvider) LoadNoteVersion(id guid.Guid, version int) (*Note, error) {
	return nil, nil
}

// UPDATE //
// Update note to given data, append version #
func (p *FileProvider) UpdateNote(id guid.Guid, updatedNote *Note) error {
	return nil
}

// Restore a Note to a specific version, append version #
func (p *FileProvider) RestoreNote(id guid.Guid, version int) (*Note, error) {
	return nil, nil
}

// DELETE //
// Delete a note (archive it)
func (p *FileProvider) DeleteNote(id guid.Guid) error {
	return nil
}
