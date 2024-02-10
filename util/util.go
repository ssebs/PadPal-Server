// util.go
package util

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
)

// ParseMDToHTML
// Uses simplecss.org CSS
func ParseMDToHTML(md []byte) []byte {
	head := []byte(`<link rel="stylesheet" href="https://cdn.simplecss.org/simple.min.css">`)

	// create markdown parser with extensions
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs | parser.NoEmptyLineBeforeBlock
	p := parser.NewWithExtensions(extensions)
	doc := p.Parse(md)

	// create HTML renderer with extensions
	htmlFlags := html.CommonFlags | html.HrefTargetBlank | html.CompletePage
	opts := html.RendererOptions{Flags: htmlFlags, Head: head}
	renderer := html.NewRenderer(opts)

	return markdown.Render(doc, renderer)
}

// GetFilenamesInDir will return a list of filenames from a directory
func GetFilenamesInDir(d string) ([]string, error) {
	filenames := make([]string, 0)
	f, err := os.Open(d)
	if err != nil {
		return filenames, err
	}

	filenames, err = f.Readdirnames(0)
	if err != nil {
		return filenames, err
	}
	return filenames, nil
}

// GetFilesInDir will return a list of files from a directory
func GetFilesInDir(d string) ([]*os.File, error) {
	files := make([]*os.File, 0)

	// Get filenames first
	filenames, err := GetFilenamesInDir(d)
	if err != nil {
		return nil, err
	}

	// Return files from filenames
	for _, file := range filenames {
		f, err := os.Open(filepath.Join(d + "/" + file))
		if err != nil {
			return nil, fmt.Errorf("failed to open %s, err:%s", file, err.Error())
		}
		files = append(files, f)
	}
	return files, nil
}

// TESTING STUFF //
// GotWantTest takes *testing.T
// Used for testing...
func GotWantTest[T comparable](got, want T, t *testing.T) {
	t.Helper()
	if got != want {
		t.Fatalf("got %v, want %v", got, want)
	}
}
