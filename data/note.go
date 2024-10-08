// note.go
package data

import (
	"encoding/json"
	"time"

	"github.com/beevik/guid"
)

// NoteBind used for form data to create a Note
// TODO: replace the name of this!
type NoteBind struct {
	Title    string `json:"title"`
	Contents string `json:"contents"`
	Author   string `json:"author"`
}

// TODO: Create NoteMeta for file_provider

// Note
// ID is a GUID
// LastUpdated is a time.Time in UTC
type Note struct {
	ID          *guid.Guid `json:"-"`
	Title       string     `json:"title"`
	Contents    string     `json:"contents"`
	Author      string     `json:"author"`
	LastUpdated time.Time  `json:"last_updated"`
}

// NewNote will create a new note from the title, author, and contents.
// LastUpdated, Version, and Active will be set by default
// Retuns a *Note
func NewNote(title, author, contents string) *Note {
	// TODO: Move Version to a db, or somewhere else..
	return &Note{
		ID:          guid.New(),
		Title:       title,
		Contents:    contents,
		Author:      author,
		LastUpdated: time.Now().UTC(),
	}
}

// NewNoteFromBind will create a new note from a mapped NoteBind
func NewNoteFromBind(nb NoteBind) *Note {
	note := NewNote(nb.Title, nb.Author, nb.Contents)
	return note
}

// MarshalJSON customizes the JSON marshaling for the Note struct.
// The ID field is represented as a string in the JSON output.
func (n *Note) MarshalJSON() ([]byte, error) {
	type Alias Note
	return json.Marshal(&struct {
		ID string `json:"id"`
		*Alias
	}{
		ID:    n.ID.String(),
		Alias: (*Alias)(n),
	})
}

// UnmarshalJSON customizes the JSON unmarshaling for the Note struct.
// The ID field is parsed from a string in the JSON input.
func (n *Note) UnmarshalJSON(data []byte) error {
	type Alias Note
	aux := &struct {
		ID string `json:"id"`
		*Alias
	}{
		Alias: (*Alias)(n),
	}

	if err := json.Unmarshal(data, aux); err != nil {
		return err
	}

	// Parse the ID string into a *guid.Guid
	parsedID, err := guid.ParseString(aux.ID)
	if err != nil {
		return err
	}

	n.ID = parsedID
	return nil
}
