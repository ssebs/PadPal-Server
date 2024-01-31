// file_provider.go - implement file based CRUDProvider
package data

import (
	"os"
	"path/filepath"
	"sync"

	"github.com/beevik/guid"
	"github.com/ssebs/padpal-server/util"
)

// FileProvider is a file directory implementation of CRUDProvider
type FileProvider struct {
	fullPath string
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
	_notes, err := initFileProvider(d)
	if err != nil {
		return nil, err
	}

	fp := &FileProvider{
		fullPath: d,
		notes:    _notes,
	}

	return fp, nil
}

// initFileProvider will load notes from existing dir +
// Checks if we're loading an existing dir or creating new git dir
func initFileProvider(d string) (map[guid.Guid]*Note, error) {
	_notes := make(map[guid.Guid]*Note)

	// Check if we're in a git repo

	// If we are, load notes from "active" folder

	// If not, init git dir,
	if err := util.InitGitDir(d); err != nil {
		return nil, err
	}

	// ...then load notes from "active folder"

	return _notes, nil
}

// CREATE //
// Save note to disk
func (p *FileProvider) SaveNote(note *Note) error {

	return nil
}

// TODO: CopyNote(id guid.Guid) error

// READ //
// List all active notes
func (p *FileProvider) ListNotes(query string) ([]*Note, error) {
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
