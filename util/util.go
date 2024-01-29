// util.go
package util

import (
	"os"
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

// TESTING STUFF //
// GotWantTest takes *testing.T
// Used for testing...
func GotWantTest[T comparable](got, want T, t *testing.T) {
	t.Helper()
	if got != want {
		t.Fatalf("got %v, want %v", got, want)
	}
}
