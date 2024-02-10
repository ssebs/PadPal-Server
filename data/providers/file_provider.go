// file_provider.go - implement file based CRUDProvider
package providers

import (
	"os"
	"path/filepath"
	"sync"

	"github.com/beevik/guid"
	"github.com/ssebs/padpal-server/data"
)

// FileProvider is a file directory implementation of CRUDProvider
type FileProvider struct {
	dirPath string
	notes   map[guid.Guid]*data.Note
	mutex   sync.RWMutex
}

// NewFileProvider
// If files exist in dir, it will load them up for CRUD'ing
// NOTE: NO Version Control exists yet!
//
// make sure to include the last / for the dir!
func NewFileProvider(dir string) (*FileProvider, error) {
	// Check dir exists
	d := filepath.Clean(dir)
	if _, err := os.Stat(d); os.IsNotExist(err) {
		return nil, err
	}

	fp := &FileProvider{
		dirPath: d,
		notes:   make(map[guid.Guid]*data.Note),
	}
	// TODO: load notes in ListNotes/LoadNotes

	return fp, nil
}

// CREATE //
// Save note to disk, under p.fullPath/active/<file>.md
func (p *FileProvider) SaveNote(note *data.Note) error {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	fullFileName := filepath.Join(p.dirPath + "/" + filepath.Clean(note.GetFilename()))

	// save md file to disk (need fullpath)
	if err := os.WriteFile(fullFileName, []byte(note.Contents), 0644); err != nil {
		return err
	}

	// TODO: git add & commit

	return nil
}

// TODO: CopyNote(id guid.Guid) error

// READ //
// LoadNotes loads notes from a query & read them into memory
func (p *FileProvider) LoadNotes(query string) ([]*data.Note, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	// Get all the Files

	// Query them

	// Add to the return value

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
func (p *FileProvider) LoadNote(id guid.Guid) (*data.Note, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	return nil, nil
}

// List all versions of a note
func (p *FileProvider) ListNoteVersions(id guid.Guid) ([]int, error) {
	return nil, nil
}

// Load note from disk by guid ID + version
func (p *FileProvider) LoadNoteVersion(id guid.Guid, version int) (*data.Note, error) {
	return nil, nil
}

// UPDATE //
// Update note to given data, append version #
func (p *FileProvider) UpdateNote(id guid.Guid, updatedNote *data.Note) error {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	return nil
}

// Restore a Note to a specific version, append version #
func (p *FileProvider) RestoreNote(id guid.Guid, version int) (*data.Note, error) {
	return nil, nil
}

// DELETE //
// Delete a note (archive it)
func (p *FileProvider) DeleteNote(id guid.Guid) error {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	return nil
}
