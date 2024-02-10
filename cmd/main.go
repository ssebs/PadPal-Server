// main.go
package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/ssebs/padpal-server/data"
	"github.com/ssebs/padpal-server/data/providers"
)

func main() {
	fmt.Println("PadPal Server")
	// TODO: Add provider param
	dir := flag.String("dir", ".", "File Provider dir path. Include trailing /")
	flag.Parse()

	// Init CRUDProvider
	fp, err := providers.NewFileProvider(*dir)
	if err != nil {
		log.Fatal(err)
	}

	// TEST: save a note
	err = fp.SaveNote(
		data.NewNote("Test-Note", "Seb", "# test-note\nGotta love me some test, *amirite*\n"),
	)
	if err != nil {
		log.Fatal(err)
	}

	// TODO: Add provider param to handleandserve
	// api.HandleAndServe("", 5000)
}
