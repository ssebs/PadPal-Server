// file_provider.go - implement file based CRUDProvider
package providers

import (
	"errors"
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

	// load notes in LoadNotes
	notes, err := fp.LoadNotes("")
	if err != nil {
		return fp, err
	}

	// add notes to the map of notes from the GUID
	for _, n := range notes {
		fp.notes[*n.ID] = n
	}

	return fp, nil
}

// CREATE //
// Save note to disk, under p.fullPath/<file>.md
// Filename will be special, see data.Note.GetFilename()
func (p *FileProvider) SaveNote(note *data.Note) error {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	fullFileName := filepath.Join(p.dirPath + "/" + filepath.Clean(note.GetFilename()))

	// save md file to disk (need fullpath)
	if err := os.WriteFile(fullFileName, []byte(note.Contents), 0644); err != nil {
		return err
	}

	p.notes[*note.ID] = note

	// TODO: git add & commit

	return nil
}

// TODO: CopyNote(id guid.Guid) error

// READ //
// LoadNotes loads notes from a query & read them into memory
// If query is empty, load all notes. Otherwise, filter by the query string
func (p *FileProvider) LoadNotes(query string) ([]*data.Note, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	notes := make([]*data.Note, 0)

	// TODO: check in p.notes before searching disk!

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

		// Query / filter them
		if n.SearchMatch(query) {
			notes = append(notes, n)
		}
	}

	// TODO: implement

	return notes, nil
}

// Load note from disk by guid ID
func (p *FileProvider) LoadNote(id guid.Guid) (*data.Note, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	// Check if the GUID is in the map
	n, ok := p.notes[id]
	if ok {
		return n, nil
	}
	// If not, check the fs
	// TODO: implement

	// If nothing exists, return filenotfound
	return nil, fmt.Errorf("%s not found in %s", id.String(), p.dirPath)
}

// List all versions of a note
func (p *FileProvider) ListNoteVersions(id guid.Guid) ([]int, error) {
	return nil, errors.New("method not implemented")
}

// Load note from disk by guid ID + version
func (p *FileProvider) LoadNoteVersion(id guid.Guid, version int) (*data.Note, error) {
	return nil, errors.New("method not implemented")
}

// UPDATE //
// Update note to given data, append version #
func (p *FileProvider) UpdateNote(id guid.Guid, updatedNote *data.Note) error {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	// TODO: see if there anything else we should do
	// ...bump version?

	fullFileName := filepath.Join(p.dirPath + "/" + filepath.Clean(updatedNote.GetFilename()))

	// save md file to disk (need fullpath)
	if err := os.WriteFile(fullFileName, []byte(updatedNote.Contents), 0644); err != nil {
		return err
	}
	return nil
}

// Restore a Note to a specific version, append version #
func (p *FileProvider) RestoreNote(id guid.Guid, version int) (*data.Note, error) {
	return nil, errors.New("method not implemented")
}

// DELETE //
// Delete a note (archive it)
func (p *FileProvider) DeleteNote(id guid.Guid) error {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	// Check if the GUID is in the map
	note, ok := p.notes[id]
	if !ok {
		return fmt.Errorf("id not found: %s", id.String())
	}
	src := filepath.Join(p.dirPath, note.GetFilename())
	dst := filepath.Join(p.dirPath, "/archive/", note.GetFilename())
	return os.Rename(src, dst)
}
