// note.go
package data

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"github.com/beevik/guid"
	"github.com/mrz1836/go-sanitize"
)

// NoteBind used for form data to create a Note
// TODO: replace the name of this!
type NoteBind struct {
	Title    string `json:"title"`
	Contents string `json:"contents"`
	Author   string `json:"author"`
}

// Note
// ID is a GUID
// LastUpdated is a time.Time in UTC
// Author MUST NOT HAVE AN UNDERSCORE CHARACTER! (see NewNoteFromFile())
type Note struct {
	ID          *guid.Guid `json:"-"`
	Title       string     `json:"title"`
	Contents    string     `json:"contents"`
	Author      string     `json:"author"`
	LastUpdated time.Time  `json:"last_updated"`
	Version     int        `json:"version"`
	Active      bool       `json:"active"`
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
		Version:     1,
		Active:      true,
	}
}

// NewNoteFromBind will create a new note from a mapped NoteBind
func NewNoteFromBind(nb NoteBind) *Note {
	note := NewNote(nb.Title, nb.Author, nb.Contents)
	return note
}

// NewNoteFromFile will take a file and parse the name/contents into a *Note
// The filename should be in the same format as GetFilename()
// e.g. "id_author_title.md"
func NewNoteFromFile(file *os.File) (*Note, error) {
	defer file.Close() // TODO: organize the open/closing of files

	filenameParts := strings.Split(filepath.Base(file.Name()), "_")
	// Parse GUID
	g, err := guid.ParseString(filenameParts[0])
	if err != nil {
		return nil, fmt.Errorf(
			"could not parse %s as a GUID from the filename %s, err:%s",
			filenameParts[0], file.Name(), err,
		)
	}

	// Parse title
	// Merge everything after, and remove the .md from the end
	t := strings.Join(filenameParts[2:], "")
	t, _ = strings.CutSuffix(t, ".md")

	// get last updated from stat
	s, _ := file.Stat()

	// Get the contents of the file
	contents, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}
	note := &Note{
		Version:     1,
		Active:      true,
		ID:          g,
		Author:      filenameParts[1],
		Title:       t,
		LastUpdated: s.ModTime(),
		Contents:    string(contents),
	}

	return note, nil
}

// GetFilename generates a filename from the GUid_author_title
// formats as "id_author_title.md"
func (n *Note) GetFilename() string {
	_author := sanitize.PathName(n.Author)
	_title := sanitize.PathName(n.Title)

	// id_author_title.md
	return fmt.Sprintf("%s_%s_%s.md", n.ID.String(), _author, _title)
}

// SearchMatch will check if a search query will match a note or not
// If query is empty, match. Supports * wildcards
func (n *Note) SearchMatch(query string) bool {
	if query == "" {
		query = "*"
	}
	query = strings.ToLower(query)

	if found, _ := regexp.MatchString(wildCardToRegexp(query), strings.ToLower(n.ID.String())); found {
		return true
	}
	if found, _ := regexp.MatchString(wildCardToRegexp(query), strings.ToLower(n.Title)); found {
		return true
	}
	if found, _ := regexp.MatchString(wildCardToRegexp(query), strings.ToLower(n.Author)); found {
		return true
	}
	if found, _ := regexp.MatchString(wildCardToRegexp(query), strings.ToLower(n.Contents)); found {
		return true
	}
	return false
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

// wildCardToRegexp converts a wildcard pattern to a regular expression pattern.
func wildCardToRegexp(pattern string) string {
	var result strings.Builder
	for i, literal := range strings.Split(pattern, "*") {

		// Replace * with .*
		if i > 0 {
			result.WriteString(".*")
		}

		// Quote any regular expression meta characters in the
		// literal text.
		result.WriteString(regexp.QuoteMeta(literal))
	}
	return result.String()
}
