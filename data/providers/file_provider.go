// file_provider.go - implement file based CRUDProvider
package providers

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"

	"github.com/beevik/guid"
	"github.com/ssebs/padpal-server/data"
	"github.com/ssebs/padpal-server/util"
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
	notes, err := fp.LoadNotes("")
	if err != nil {
		return fp, err
	}
	fmt.Println("Loading all notes:")
	fmt.Println(notes)
	// add notes to the map of notes from the GUID
	// for range...

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
	notes := make([]*data.Note, 0)

	// Get all the Files
	files, err := util.GetFilesInDir(p.dirPath)
	if err != nil {
		return nil, fmt.Errorf("failed to load notes in %s, err:%s", p.dirPath, err.Error())
	}
	// Parse the file and add to notes
	for _, f := range files {
		n, err := data.NewNoteFromFile(f)
		if err != nil {
			return nil, fmt.Errorf("failed to parse note from file %s, err:%s", f.Name(), err.Error())
		}
		notes = append(notes, n)
	}

	// Query / filter them

	return notes, nil
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
